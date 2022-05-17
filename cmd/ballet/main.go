package main

import (
	"flag"
	"fmt"
	"os"

	"chalkan.github.com/internal/ballet"
	nats_cli "chalkan.github.com/internal/nats"
	"github.com/go-kit/log"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		httpPort = fs.String("http_port", "8080", "application http port")
	)

	flag.Usage = fs.Usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	logger := createLogger(*httpPort)
	nc, err := nats_cli.NewNATS("", true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	svc := ballet.NewLoggerMW(logger, ballet.NewService())
	ballet.SubscriberTransport(svc, nc.GetConn())

}

func createLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
