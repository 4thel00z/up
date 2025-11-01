package main

import (
	"errors"
	"os"
	"runtime"

	"github.com/stripe/stripe-go"
	"github.com/tj/go/env"
	"github.com/tj/go/term"

	// commands
	_ "github.com/4thel00z/up/internal/cli/build"
	_ "github.com/4thel00z/up/internal/cli/config"
	_ "github.com/4thel00z/up/internal/cli/deploy"
	_ "github.com/4thel00z/up/internal/cli/disable-stats"
	_ "github.com/4thel00z/up/internal/cli/docs"
	_ "github.com/4thel00z/up/internal/cli/domains"
	_ "github.com/4thel00z/up/internal/cli/logs"
	_ "github.com/4thel00z/up/internal/cli/metrics"
	_ "github.com/4thel00z/up/internal/cli/prune"
	_ "github.com/4thel00z/up/internal/cli/run"
	_ "github.com/4thel00z/up/internal/cli/stack"
	_ "github.com/4thel00z/up/internal/cli/start"
	_ "github.com/4thel00z/up/internal/cli/team"
	_ "github.com/4thel00z/up/internal/cli/upgrade"
	_ "github.com/4thel00z/up/internal/cli/url"
	_ "github.com/4thel00z/up/internal/cli/version"

	"github.com/4thel00z/up/internal/cli/app"
	"github.com/4thel00z/up/internal/signal"
	"github.com/4thel00z/up/internal/stats"
	"github.com/4thel00z/up/internal/util"
)

var version = "master"

func main() {
	signal.Add(reset)
	stripe.Key = env.GetDefault("STRIPE_KEY", "pk_live_23pGrHcZ2QpfX525XYmiyzmx")
	stripe.LogLevel = 0

	err := run()

	if err == nil {
		return
	}

	term.ShowCursor()

	switch {
	case util.IsNoCredentials(err):
		util.Fatal(errors.New("Cannot find credentials, visit https://apex.sh/docs/up/credentials/ for help."))
	default:
		util.Fatal(err)
	}
}

// run the cli.
func run() error {
	stats.SetProperties(map[string]interface{}{
		"os":      runtime.GOOS,
		"arch":    runtime.GOARCH,
		"version": version,
		"ci":      os.Getenv("CI") == "true" || os.Getenv("CI") == "1",
	})

	return app.Run(version)
}

// reset cursor.
func reset() error {
	term.ShowCursor()
	println()
	return nil
}
