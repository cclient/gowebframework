package main

import (
	//	"github.com/Sirupsen/logrus"
	//	"api/client"
	//	flag "server/pkg/mflag"
	//	"server/pkg/reexec"
	//	"server/pkg/term"
	//	"crypto/tls"
	"fmt"
	//	"io"
	"os"
	//	"path/filepath"
	//	"strings"
	//	"time"
	"github.com/Sirupsen/logrus"
	apiserver "server/api"
	//	"github.com/docker/distribution/uuid"
	tool "server/common/tool"
	"server/pkg/pidfile"
	"server/pkg/signal"
	//	"server/pkg/system"
	//	"server/pkg/timeutils"
	//	"server/pkg/tlsconfig"
	//	"utils"
	"flag"
)

var PUBFILEPATH = flag.String("d", "public", "public file root path")
var HOST = flag.String("h", "0.0.0.0", "host")
var PORT = flag.String("p", "9900", "port")
var PIDFILE = flag.String("f", "apiserver.pid", "pidfile")

//var SCMD = flag.String("s", "start", "start/stop/restart")
var MGOPORT = flag.String("mgop", "27017", "mongodb port")

func main() {
	flag.Parse()
	tool.Port = *MGOPORT
	if PUBFILEPATH != nil {
		fmt.Println("dir =", *PUBFILEPATH, " host=", *HOST, " port="+*PORT, " mgop="+*MGOPORT)
		serverConfig := &apiserver.Config{
			Logging: true,
			Host:    *HOST,
			Port:    *PORT,
			//	//cdpadd
			//	PubFilePath string
			Version:     "version 1.0",
			PubFilePath: *PUBFILEPATH,
			Pidfile:     *PIDFILE,
		}
		handleGlobalDaemonFlag(serverConfig)
	} else {
		os.Exit(0)
	}
	//	else{
	//		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	//		}
	//	if reexec.Init() {
	//		return
	//	}

	// Set terminal emulation based on platform as required.
	//	stdin, stdout, stderr := term.StdStreams()
	//
	//	logrus.SetOutput(stderr)

	//	fmt.Println("1")

	// TODO: remove once `-d` is retired

	//	clientCli := client.NewDockerCli(stdin, stdout, stderr, clientFlags)
	//
	//	c := cli.New(clientCli, daemonCli)
	//	if err := c.Run(flag.Args()...); err != nil {
	//		if sterr, ok := err.(cli.StatusError); ok {
	//			if sterr.Status != "" {
	//				fmt.Fprintln(os.Stderr, sterr.Status)
	//				os.Exit(1)
	//			}
	//			os.Exit(sterr.StatusCode)
	//		}
	//		fmt.Fprintln(os.Stderr, err)
	//		os.Exit(1)
	//	}
}

//const daemonUsage = "       docker daemon [ --help | ... ]\n"

//var (
//	daemonCli cli.Handler = NewDaemonCli()
//)

// TODO: remove once `-d` is retired
func handleGlobalDaemonFlag(serverConfig *apiserver.Config) {
	CmdDaemon(serverConfig)
	os.Exit(0)
}

func presentInHelp(usage string) string { return usage }
func absentFromHelp(string) string      { return "" }

// CmdDaemon is the daemon command, called the raw arguments after `docker daemon`.
func CmdDaemon(serverConfig *apiserver.Config) error {
	// warn from uuid package when running the daemon
	//	uuid.Loggerf = logrus.Warnf

	//	if *flDaemon {
	//		// allow legacy forms `docker -D -d` and `docker -d -D`
	//		logrus.Warn("please use 'docker daemon' instead.")
	//	} else if !commonFlags.FlagSet.IsEmpty() || !clientFlags.FlagSet.IsEmpty() {
	//		// deny `docker -D daemon`
	//		illegalFlag := getGlobalFlag()
	//		fmt.Fprintf(os.Stderr, "invalid flag '-%s'.\nSee 'docker daemon --help'.\n", illegalFlag.Names[0])
	//		os.Exit(1)
	//	} else {
	//		// allow new form `docker daemon -D`
	//		flag.Merge(daemonFlags, commonFlags.FlagSet)
	//	}

	//	daemonFlags.ParseFlags(args, true)
	//	commonFlags.PostParse()

	//	if utils.ExperimentalBuild() {
	//		logrus.Warn("Running experimental build")
	//	}

	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02T15:04:05.000000000Z07:00"})

	//	if err := setDefaultUmask(); err != nil {
	//		logrus.Fatalf("Failed to set umask: %v", err)
	//	}
	//
	//	if len(cli.LogConfig.Config) > 0 {
	//		if err := logger.ValidateLogOpts(cli.LogConfig.Type, cli.LogConfig.Config); err != nil {
	//			logrus.Fatalf("Failed to set log opts: %v", err)
	//		}
	//	}

	var pfile *pidfile.PIDFile
	if serverConfig.Pidfile != "" {
		pf, err := pidfile.New(serverConfig.Pidfile)
		if err != nil {
			logrus.Fatalf("Error starting daemon: %v", err)
		}
		pfile = pf
		defer func() {
			if err := pfile.Remove(); err != nil {
				logrus.Error(err)
			}
		}()
	}

	//	serverConfig := &apiserver.Config{
	//		Logging: true,
	//
	//		//		  Host       "
	//		//    Port       int
	//		//	//cdpadd
	//		//	PubFilePath string
	//
	//		Version: "version 1.0",
	//
	//		//		PubFilePath:PUBFILEPATH,
	//	}

	//	serverConfig = setPlatformServerConfig(serverConfig, cli.Config)

	//	if commonFlags.TLSOptions != nil {
	//		if !commonFlags.TLSOptions.InsecureSkipVerify {
	//			// server requires and verifies client's certificate
	//			commonFlags.TLSOptions.ClientAuth = tls.RequireAndVerifyClientCert
	//		}
	//		tlsConfig, err := tlsconfig.Server(*commonFlags.TLSOptions)
	//		if err != nil {
	//			logrus.Fatal(err)
	//		}
	//		serverConfig.TLSConfig = tlsConfig
	//	}

	//	if len(commonFlags.Hosts) == 0 {
	//		commonFlags.Hosts = make([]string, 1)
	//	}
	//	for i := 0; i < len(commonFlags.Hosts); i++ {
	//		var err error
	//		if commonFlags.Hosts[i], err = opts.ParseHost(commonFlags.Hosts[i]); err != nil {
	//			logrus.Fatalf("error parsing -H %s : %v", commonFlags.Hosts[i], err)
	//		}
	//	}
	//	for _, protoAddr := range commonFlags.Hosts {
	//		protoAddrParts := strings.SplitN(protoAddr, "://", 2)
	//		if len(protoAddrParts) != 2 {
	//			logrus.Fatalf("bad format %s, expected PROTO://ADDR", protoAddr)
	//		}
	//		serverConfig.Addrs = append(serverConfig.Addrs, apiserver.Addr{Proto: protoAddrParts[0], Addr: protoAddrParts[1]})
	//	}

	api, err := apiserver.New(serverConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	// The serve API routine never exits unless an error occurs
	// We need to start it as a goroutine and wait on it so
	// daemon doesn't exit
	// All servers must be protected with some mechanism (systemd socket, listenbuffer)
	// which prevents real handling of request until routes will be set.
	serveAPIWait := make(chan error)
	go func() {
		if err := api.ServeAPI(); err != nil {
			logrus.Errorf("ServeAPI error: %v", err)
			serveAPIWait <- err
			return
		}
		serveAPIWait <- nil
	}()

	//	if err := migrateKey(); err != nil {
	//		logrus.Fatal(err)
	//	}
	//	cli.TrustKeyPath = commonFlags.TrustKey

	//	registryService := registry.NewService(cli.registryOptions)
	//	d, err := daemon.NewDaemon(cli.Config, registryService)
	//	if err != nil {
	//		if pfile != nil {
	//			if err := pfile.Remove(); err != nil {
	//				logrus.Error(err)
	//			}
	//		}
	//		logrus.Fatalf("Error starting daemon: %v", err)
	//	}

	logrus.Info("Daemon has completed initialization")

	//	logrus.WithFields(logrus.Fields{
	//		"version":     dockerversion.VERSION,
	//		"commit":      dockerversion.GITCOMMIT,
	//		"execdriver":  d.ExecutionDriver().Name(),
	//		"graphdriver": d.GraphDriver().String(),
	//	}).Info("Docker daemon")

	//"/Users/cdpmac/Documents/workspace/c/audit/src/public"
	api.InitRouters()

	signal.Trap(func() {
		api.Close()
		<-serveAPIWait
		//		shutdownDaemon(d, 15)
		if pfile != nil {
			if err := pfile.Remove(); err != nil {
				logrus.Error(err)
			}
		}
	})

	// after the daemon is done setting up we can tell the api to start
	// accepting connections with specified daemon
	//	notifySystem()

	api.AcceptConnections()

	//	 Daemon is fully initialized and handling API traffic
	//	 Wait for serve API to complete
	errAPI := <-serveAPIWait
	//	shutdownDaemon(d, 15)
	if errAPI != nil {
		if pfile != nil {
			if err := pfile.Remove(); err != nil {
				logrus.Error(err)
			}
		}
		logrus.Fatalf("Shutting down due to ServeAPI error: %v", errAPI)
	}
	return nil
}

// shutdownDaemon just wraps daemon.Shutdown() to handle a timeout in case
// d.Shutdown() is waiting too long to kill container or worst it's
// blocked there
//func shutdownDaemon(d *daemon.Daemon, timeout time.Duration) {
//	ch := make(chan struct{})
//	go func() {
//		d.Shutdown()
//		close(ch)
//	}()
//	select {
//	case <-ch:
//		logrus.Debug("Clean shutdown succeeded")
//	case <-time.After(timeout * time.Second):
//		logrus.Error("Force shutdown daemon")
//	}
//}

// Shutdown stops the daemon.
//func (daemon *Daemon) Shutdown() error {
//	daemon.shutdown = true
//	if daemon.containers != nil {
//		group := sync.WaitGroup{}
//		logrus.Debug("starting clean shutdown of all containers...")
//		for _, container := range daemon.List() {
//			c := container
//			if c.IsRunning() {
//				logrus.Debugf("stopping %s", c.ID)
//				group.Add(1)
//
//				go func() {
//					defer group.Done()
//					// TODO(windows): Handle docker restart with paused containers
//					if c.isPaused() {
//						// To terminate a process in freezer cgroup, we should send
//						// SIGTERM to this process then unfreeze it, and the process will
//						// force to terminate immediately.
//						logrus.Debugf("Found container %s is paused, sending SIGTERM before unpause it", c.ID)
//						sig, ok := signal.SignalMap["TERM"]
//						if !ok {
//							logrus.Warnf("System does not support SIGTERM")
//							return
//						}
//						if err := daemon.kill(c, int(sig)); err != nil {
//							logrus.Debugf("sending SIGTERM to container %s with error: %v", c.ID, err)
//							return
//						}
//						if err := c.unpause(); err != nil {
//							logrus.Debugf("Failed to unpause container %s with error: %v", c.ID, err)
//							return
//						}
//						if _, err := c.WaitStop(10 * time.Second); err != nil {
//							logrus.Debugf("container %s failed to exit in 10 second of SIGTERM, sending SIGKILL to force", c.ID)
//							sig, ok := signal.SignalMap["KILL"]
//							if !ok {
//								logrus.Warnf("System does not support SIGKILL")
//								return
//							}
//							daemon.kill(c, int(sig))
//						}
//					} else {
//						// If container failed to exit in 10 seconds of SIGTERM, then using the force
//						if err := c.Stop(10); err != nil {
//							logrus.Errorf("Stop container %s with error: %v", c.ID, err)
//						}
//					}
//					c.WaitStop(-1 * time.Second)
//					logrus.Debugf("container stopped %s", c.ID)
//				}()
//			}
//		}
//		group.Wait()
//
//		// trigger libnetwork Stop only if it's initialized
//		if daemon.netController != nil {
//			daemon.netController.Stop()
//		}
//	}
//
//	if daemon.containerGraphDB != nil {
//		if err := daemon.containerGraphDB.Close(); err != nil {
//			logrus.Errorf("Error during container graph.Close(): %v", err)
//		}
//	}
//
//	if daemon.driver != nil {
//		if err := daemon.driver.Cleanup(); err != nil {
//			logrus.Errorf("Error during graph storage driver.Cleanup(): %v", err)
//		}
//	}
//
//	if err := daemon.cleanupMounts(); err != nil {
//		return err
//	}
//
//	return nil
//}
