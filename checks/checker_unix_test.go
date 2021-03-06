// +build linux darwin freebsd netbsd

package checks

import (
	"testing"
	"time"

	"github.com/mackerelio/mackerel-agent/config"
	"github.com/mackerelio/mackerel-agent/util"
)

func TestChecker_CheckTimeout(t *testing.T) {
	origTimeoutDur := util.TimeoutDuration
	defer func() {
		util.TimeoutDuration = origTimeoutDur
	}()
	util.TimeoutDuration = 1 * time.Second

	checkerTimeout := Checker{
		Config: config.PluginConfig{
			Command: "sleep 2",
		},
	}

	{
		report, err := checkerTimeout.Check()
		if err != nil {
			t.Errorf("err should be nil: %v", err)
		}
		if report.Status != StatusUnknown {
			t.Errorf("status should be UNKNOWN: %v", report.Status)
		}
		if report.Message != "command timed out" {
			t.Errorf("wrong message: %q", report.Message)
		}
	}
}
