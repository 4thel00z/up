package main

import (
	"os"
	"time"

	"github.com/apex/go-apex"
	"github.com/apex/log"
	"github.com/apex/log/handlers/json"

	"github.com/4thel00z/up"
	"github.com/4thel00z/up/handler"
	"github.com/4thel00z/up/internal/logs"
	"github.com/4thel00z/up/internal/proxy"
	"github.com/4thel00z/up/internal/util"
	"github.com/4thel00z/up/platform/aws/runtime"
)

func main() {
	start := time.Now()
	stage := os.Getenv("UP_STAGE")

	// setup logging
	log.SetHandler(json.Default)
	if s := os.Getenv("LOG_LEVEL"); s != "" {
		log.SetLevelFromString(s)
	}

	log.Log = log.WithFields(logs.Fields())
	log.Info("initializing")

	// read config
	c, err := up.ReadConfig("up.json")
	if err != nil {
		log.Fatalf("error reading config: %s", err)
	}

	ctx := log.WithFields(log.Fields{
		"name": c.Name,
		"type": c.Type,
	})

	// init project
	p := runtime.New(c)

	// init runtime
	if err := p.Init(stage); err != nil {
		ctx.Fatalf("error initializing: %s", err)
	}

	// overrides
	if err := c.Override(stage); err != nil {
		ctx.Fatalf("error overriding: %s", err)
	}

	// create handler
	h, err := handler.FromConfig(c)
	if err != nil {
		ctx.Fatalf("error creating handler: %s", err)
	}

	// init handler
	h, err = handler.New(c, h)
	if err != nil {
		ctx.Fatalf("error initializing handler: %s", err)
	}

	// serve
	log.WithField("duration", util.MillisecondsSince(start)).Info("initialized")
	apex.Handle(proxy.NewHandler(h))
}
