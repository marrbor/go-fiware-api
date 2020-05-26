// orion server for testing.
package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/marrbor/go-fiware-api/orion"
	"github.com/mattn/go-pipeline"
)

const (
	DefaultOrionHost            = "localhost"
	DefaultOrionPort            = 1026
	DefaultWaitForServerStartup = 5 // second
	DefaultWaitForServerStop    = 5 // second

	TestProject = "orion-test"
)

var (
	Host string
	Port int
)

func init() {
	Host = DefaultOrionHost
	Port = DefaultOrionPort
}

func StartTestServer(t *testing.T) error {
	// check server has been already started.
	stop, err := isTestServerStop()
	if err != nil {
		return err
	}

	if stop {
		// start server
		cmd := exec.Command("docker-compose", fmt.Sprintf("-p %s", TestProject), "up", "-d")
		out, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		t.Log(string(out))
	}

	// waiting for server has been activated for X seconds.
	baseURL := fmt.Sprintf("http://%s:%d", Host, Port)
	a := orion.NewAccessor(baseURL)
	for i := 0; i < DefaultWaitForServerStartup; i++ {
		v, err := a.GetVersion()
		if err == nil && v != nil {
			t.Logf("server started up. within %d second(s)", i)
			t.Logf("version:%+v", v)
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("cannot detect server start")
}

// isTestServerStop returns whether test server is stopped.
// stop means `docker ps -q` return no outputs.
func isTestServerStop() (bool, error) {
	out, err := pipeline.CombinedOutput(
		[]string{"docker", "ps", "-q"},
		[]string{"wc", "-l"},
	)
	if err != nil {
		return false, err
	}
	s := strings.TrimSpace(string(out))
	return s == "0", nil
}

func StopTestServer(t *testing.T) error {
	// check server has been already stopped.
	stop, err := isTestServerStop()
	if err != nil {
		return err
	}
	if stop {
		t.Log("server already stopped. nothing to do.")
		return nil // already stopped.
	}

	cmd := exec.Command("docker-compose", fmt.Sprintf("-p %s", TestProject), "down", "--remove-orphans")
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// check server status for X second.
	for i := 0; i < DefaultWaitForServerStop; i++ {
		stop, err := isTestServerStop()
		if err != nil {
			return err
		}
		if stop {
			t.Logf("detect server stop within %d second(s)", i)
			return nil // ok
		}
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("cannot detect server stop")
}
