package info

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
	"strings"

	"github.com/takama/router"
)

// ServiceInfo defines HTTP API response giving service information
type ServiceInfo struct {
	Host    string       `json:"host"`
	Version string       `json:"version"`
	Repo    string       `json:"repo"`
	Commit  string       `json:"commit"`
	Build   string       `json:"build"`
	Runtime *RuntimeInfo `json:"runtime"`
}

// RuntimeInfo defines runtime part of service information
type RuntimeInfo struct {
	Compiler   string `json:"compilier"`
	CPU        int    `json:"cpu"`
	Memory     string `json:"memory"`
	Goroutines int    `json:"goroutines"`
	Uptime     string `json:"uptime"`
}

var startTime time.Time

func init() {
	startTime = time.Now()
}

func uptime() time.Duration {
    return time.Since(startTime)
}

func shortDur(d time.Duration) string {
    s := d.String()
    if strings.HasSuffix(s, "m0s") {
        s = s[:len(s)-2]
    }
    if strings.HasSuffix(s, "h0m") {
        s = s[:len(s)-2]
    }
    return s
}

// Handler provides JSON API response giving service information
func Handler(version, repo, commit, build string) router.Handle {
	return func(c *router.Control) {
		host, _ := os.Hostname()
		m := new(runtime.MemStats)
		runtime.ReadMemStats(m)

		rt := &RuntimeInfo{
			Compiler:   runtime.Version(),
			CPU:        runtime.NumCPU(),
			Memory:     fmt.Sprintf("%.2fMB", float64(m.Alloc)/(1<<(10*2))),
			Goroutines: runtime.NumGoroutine(),
			Uptime:     shortDur(uptime()),
		}

		info := ServiceInfo{
			Host:    host,
			Runtime: rt,
			Version: version,
			Repo:    repo,
			Commit:  commit,
			Build:   build,
		}

		c.Code(http.StatusOK).Body(info)
	}
}
