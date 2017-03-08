package api

import (
	"crypto/tls"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net"
	"net/http"
	"os"
	"server/api/router"
	"server/api/v1/account"
	"server/api/v1/local"
	"server/common/httputils"
	"server/pkg/libnetwork/portallocator"
	"server/pkg/listenbuffer"
	"strconv"
	"strings"
	"time"
	"utils"
)

// versionMatcher defines a variable matcher to be parsed by the router
// when a request is about to be served.
const versionMatcher = "/v{version:[0-9.]+}"

// Config provides the configuration for the API server
type Config struct {
	Logging     bool
	EnableCors  bool
	CorsHeaders string
	Version     string
	SocketGroup string
	TLSConfig   *tls.Config
	//	Addrs       []Addr
	Host        string
	Port        string
	PubFilePath string
	Pidfile     string
}

// Server contains instance details for the server
type Server struct {
	cfg     *Config
	start   chan struct{}
	servers []*HTTPServer
	routers []router.Router
}

// Addr contains string representation of address and its protocol (tcp, unix...).
type Addr struct {
	Proto string
	Addr  string
}

// New returns a new instance of the server based on the specified configuration.
// It allocates resources which will be needed for ServeAPI(ports, unix-sockets).
func New(cfg *Config) (*Server, error) {
	s := &Server{
		cfg:   cfg,
		start: make(chan struct{}),
	}
	//	for _, addr := range cfg.Addrs {
	//		//		srv, err := s.newServer(addr.Proto, addr.Addr)
	//		srv, err := s.newServer("tcp", "0.0.0.0:9900")
	//		if err != nil {
	//			return nil, err
	//		}
	//		logrus.Debugf("Server created for HTTP on %s (%s)", addr.Proto, addr.Addr)
	//		s.servers = append(s.servers, srv...)
	//	}
	//	srv, err := s.newServer("tcp", "0.0.0.0:8990")
	srv, err := s.newServer("tcp", s.cfg.Host+":"+s.cfg.Port)
	if err != nil {
		return nil, err
	}
	logrus.Debugf("Server created for HTTP on %s (%s)", "tcp", s.cfg.Host+":"+s.cfg.Port)
	s.servers = append(s.servers, srv...)

	return s, nil
}

// Close closes servers and thus stop receiving requests
func (s *Server) Close() {
	for _, srv := range s.servers {
		if err := srv.Close(); err != nil {
			logrus.Error(err)
		}
	}
}

// ServeAPI loops through all initialized servers and spawns goroutine
// with Serve() method for each.
func (s *Server) ServeAPI() error {
	var chErrors = make(chan error, len(s.servers))
	for _, srv := range s.servers {
		go func(srv *HTTPServer) {
			var err error
			logrus.Infof("API listen on %s", srv.l.Addr())
			if err = srv.Serve(); err != nil && strings.Contains(err.Error(), "use of closed network connection") {
				err = nil
			}
			chErrors <- err
		}(srv)
	}
	for i := 0; i < len(s.servers); i++ {
		err := <-chErrors
		if err != nil {
			return err
		}
	}
	return nil
}

// HTTPServer contains an instance of http server and the listener.
// srv *http.Server, contains configuration to create a http server and a mux router with all api end points.
// l   net.Listener, is a TCP or Socket listener that dispatches incoming request to the router.
type HTTPServer struct {
	srv *http.Server
	l   net.Listener
}

// Serve starts listening for inbound requests.
func (s *HTTPServer) Serve() error {
	return s.srv.Serve(s.l)
}

// Close closes the HTTPServer from listening for the inbound requests.
func (s *HTTPServer) Close() error {
	return s.l.Close()
}

func writeCorsHeaders(w http.ResponseWriter, r *http.Request, corsHeaders string) {
	logrus.Debugf("CORS header is enabled and set to: %s", corsHeaders)
	w.Header().Add("Access-Control-Allow-Origin", corsHeaders)
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth")
	w.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
}

func (s *Server) initTCPSocket(addr string) (l net.Listener, err error) {
	//	if s.cfg.TLSConfig == nil || s.cfg.TLSConfig.ClientAuth != tls.RequireAndVerifyClientCert {
	//		logrus.Warn("/!\\ DON'T BIND ON ANY IP ADDRESS WITHOUT setting -tlsverify IF YOU DON'T KNOW WHAT YOU'RE DOING /!\\")
	//	}
	//	if l, err = sockets.NewTCPSocket(addr, s.cfg.TLSConfig, s.start); err != nil {
	//	if l, err = NewTCPSocket(addr, s.cfg.TLSConfig, s.start); err != nil {
	//		return nil, err
	//	}
	if l, err = NewTCPSocketNoTls(addr, s.start); err != nil {
		return nil, err
	}

	if err := allocateDaemonPort(addr); err != nil {
		return nil, err
	}
	return
}

func (s *Server) makeHTTPHandler(handler httputils.APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log the handler call
		logrus.Debugf("Calling %s %s", r.Method, r.URL.Path)
		// Define the context that we'll pass around to share info
		// like the docker-request-id.
		//
		// The 'context' will be used for global data that should
		// apply to all requests. Data that is specific to the
		// immediate function being called should still be passed
		// as 'args' on the function call.
		ctx := context.Background()
		handlerFunc := s.handleWithGlobalMiddlewares(handler)

		vars := mux.Vars(r)
		if vars == nil {
			vars = make(map[string]string)
		}

		if err := handlerFunc(ctx, w, r, vars); err != nil {
			logrus.Errorf("Handler for %s %s returned error: %s", r.Method, r.URL.Path, utils.GetErrorMessage(err))
			httputils.WriteError(w, err)
		}
	}
}

// InitRouters initializes a list of routers for the server.
// Sets those routers as Handler for each server.
func (s *Server) InitRouters() {
	s.addRouter(local.NewRouter())
	s.addRouter(account.NewRouter())
	for _, srv := range s.servers {
		newmux := http.NewServeMux()
		filemux := http.FileServer(http.Dir(s.cfg.PubFilePath))
		apimux := s.CreateMux()
		newmux.Handle("/", filemux)
		newmux.Handle("/api/", apimux)
		srv.srv.Handler = newmux
	}
}

// addRouter adds a new router to the server.
func (s *Server) addRouter(r router.Router) {
	s.routers = append(s.routers, r)
}

// CreateMux initializes the main router the server uses.
// we keep enableCors just for legacy usage, need to be removed in the future
func (s *Server) CreateMux() *mux.Router {
	m := mux.NewRouter()
	if os.Getenv("DEBUG") != "" {
		profilerSetup(m, "/debug/")
	}
	fmt.Println("Registering routers")
	logrus.Debugf("Registering routers")
	for _, apiRouter := range s.routers {
		for _, r := range apiRouter.Routes() {
			f := s.makeHTTPHandler(r.Handler())
			logrus.Debugf("Registering %s, %s", r.Method(), r.Path())
			//			m.Path(versionMatcher + r.Path()).Methods(r.Method()).Handler(f)
			m.Path(r.Path()).Methods(r.Method()).Handler(f)
		}
	}
	return m
}

// AcceptConnections allows clients to connect to the API server.
// Referenced Daemon is notified about this server, and waits for the
// daemon acknowledgement before the incoming connections are accepted.
func (s *Server) AcceptConnections() {
	fmt.Println("AcceptConnections")
	// close the lock so the listeners start accepting connections
	select {
	case <-s.start:
	default:
		close(s.start)
	}
}

// newServer sets up the required HTTPServers and does protocol specific checking.
// newServer does not set any muxers, you should set it later to Handler field
func (s *Server) newServer(proto, addr string) ([]*HTTPServer, error) {
	var (
		//err error
		ls []net.Listener
	)
	switch proto {
	case "tcp":
		l, err := s.initTCPSocket(addr)
		if err != nil {
			return nil, err
		}
		ls = append(ls, l)
	default:
		return nil, fmt.Errorf("Invalid protocol format: %q", proto)
	}

	var res []*HTTPServer
	for _, l := range ls {
		httpsrv := &http.Server{
			Addr: addr,
		}
		res = append(res, &HTTPServer{
			httpsrv,
			l,
		})
	}
	return res, nil
}

func allocateDaemonPort(addr string) error {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return err
	}

	intPort, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	var hostIPs []net.IP
	if parsedIP := net.ParseIP(host); parsedIP != nil {
		hostIPs = append(hostIPs, parsedIP)
	} else if hostIPs, err = net.LookupIP(host); err != nil {
		return fmt.Errorf("failed to lookup %s address in host specification", host)
	}

	pa := portallocator.Get()
	for _, hostIP := range hostIPs {
		if _, err := pa.RequestPort(hostIP, "tcp", intPort); err != nil {
			return fmt.Errorf("failed to allocate daemon listening port %d (err: %v)", intPort, err)
		}
	}
	return nil
}

// NewTCPSocket creates a TCP socket listener with the specified address and
// and the specified tls configuration. If TLSConfig is set, will encapsulate the
// TCP listener inside a TLS one.
// The channel passed is used to activate the listenbuffer when the caller is ready
// to accept connections.
func NewTCPSocket(addr string, tlsConfig *tls.Config, activate <-chan struct{}) (net.Listener, error) {
	l, err := listenbuffer.NewListenBuffer("tcp", addr, activate)
	if err != nil {
		return nil, err
	}
	if tlsConfig != nil {
		tlsConfig.NextProtos = []string{"http/1.1"}
		l = tls.NewListener(l, tlsConfig)
	}
	return l, nil
}
func NewTCPSocketNoTls(addr string, activate <-chan struct{}) (net.Listener, error) {
	fmt.Println(addr)
	l, err := listenbuffer.NewListenBuffer("tcp", addr, activate)
	fmt.Println(l)

	if err != nil {
		return nil, err
	}
	//	if tlsConfig != nil {
	//		tlsConfig.NextProtos = []string{"http/1.1"}
	//		l = tls.NewListener(l, tlsConfig)
	//	}
	//	ln, err := Listen("tcp", "127.0.0.1:0")
	//	if err != nil {
	//		if ln, err = Listen("tcp6", "[::1]:0"); err != nil {
	//			t.Fatalf("ListenTCP on :0: %v", err)
	//		}
	//	}
	return l, nil
}

// ConfigureTCPTransport configures the specified Transport according to the
// specified proto and addr.
// If the proto is unix (using a unix socket to communicate) the compression
// is disabled.
func ConfigureTCPTransport(tr *http.Transport, proto, addr string) {
	// Why 32? See https://github.com/docker/docker/pull/8035.
	timeout := 32 * time.Second
	if proto == "unix" {
		// No need for compression in local communications.
		tr.DisableCompression = true
		tr.Dial = func(_, _ string) (net.Conn, error) {
			return net.DialTimeout(proto, addr, timeout)
		}
	} else {
		tr.Proxy = http.ProxyFromEnvironment
		tr.Dial = (&net.Dialer{Timeout: timeout}).Dial
	}
}
