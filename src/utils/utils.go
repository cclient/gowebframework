package utils

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"errcode"
)

// SelfPath figures out the absolute path of our own binary (if it's still around).
func SelfPath() string {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		if execErr, ok := err.(*exec.Error); ok && os.IsNotExist(execErr.Err) {
			return ""
		}
		panic(err)
	}
	path, err = filepath.Abs(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		panic(err)
	}
	return path
}

func dockerInitSha1(target string) string {
	f, err := os.Open(target)
	if err != nil {
		return ""
	}
	defer f.Close()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}


// DockerInitPath figures out the path of our dockerinit (which may be SelfPath())


var globalTestID string


// GetCallerName introspects the call stack and returns the name of the
// function `depth` levels down in the stack.
func GetCallerName(depth int) string {
	// Use the caller function name as a prefix.
	// This helps trace temp directories back to their test.
	pc, _, _, _ := runtime.Caller(depth + 1)
	callerLongName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(callerLongName, ".")
	callerShortName := parts[len(parts)-1]
	return callerShortName
}

// ReplaceOrAppendEnvValues returns the defaults with the overrides either
// replaced by env key or appended to the list
func ReplaceOrAppendEnvValues(defaults, overrides []string) []string {
	cache := make(map[string]int, len(defaults))
	for i, e := range defaults {
		parts := strings.SplitN(e, "=", 2)
		cache[parts[0]] = i
	}

	for _, value := range overrides {
		// Values w/o = means they want this env to be removed/unset.
		if !strings.Contains(value, "=") {
			if i, exists := cache[value]; exists {
				defaults[i] = "" // Used to indicate it should be removed
			}
			continue
		}

		// Just do a normal set/update
		parts := strings.SplitN(value, "=", 2)
		if i, exists := cache[parts[0]]; exists {
			defaults[i] = value
		} else {
			defaults = append(defaults, value)
		}
	}

	// Now remove all entries that we want to "unset"
	for i := 0; i < len(defaults); i++ {
		if defaults[i] == "" {
			defaults = append(defaults[:i], defaults[i+1:]...)
			i--
		}
	}

	return defaults
}

// ReadDockerIgnore reads a .dockerignore file and returns the list of file patterns
// to ignore. Note this will trim whitespace from each line as well
// as use GO's "clean" func to get the shortest/cleanest path for each.
func ReadDockerIgnore(reader io.ReadCloser) ([]string, error) {
	if reader == nil {
		return nil, nil
	}
	defer reader.Close()
	scanner := bufio.NewScanner(reader)
	var excludes []string

	for scanner.Scan() {
		pattern := strings.TrimSpace(scanner.Text())
		if pattern == "" {
			continue
		}
		pattern = filepath.Clean(pattern)
		excludes = append(excludes, pattern)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading .dockerignore: %v", err)
	}
	return excludes, nil
}

// ImageReference combines `repo` and `ref` and returns a string representing
// the combination. If `ref` is a digest (meaning it's of the form
// <algorithm>:<digest>, the returned string is <repo>@<ref>. Otherwise,
// ref is assumed to be a tag, and the returned string is <repo>:<tag>.
func ImageReference(repo, ref string) string {
	if DigestReference(ref) {
		return repo + "@" + ref
	}
	return repo + ":" + ref
}

// DigestReference returns true if ref is a digest reference; i.e. if it
// is of the form <algorithm>:<digest>.
func DigestReference(ref string) bool {
	return strings.Contains(ref, ":")
}

// GetErrorMessage returns the human readable message associated with
// the passed-in error. In some cases the default Error() func returns
// something that is less than useful so based on its types this func
// will go and get a better piece of text.
func GetErrorMessage(err error) string {
	switch err.(type) {
	case errcode.Error:
		e, _ := err.(errcode.Error)
		return e.Message

	case errcode.ErrorCode:
		ec, _ := err.(errcode.ErrorCode)
		return ec.Message()

	default:
		return err.Error()
	}
}
