package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	nats_cli "chalkan.github.com/internal/nats"
	"chalkan.github.com/internal/sloth"

	"github.com/go-kit/log"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		httpPort    = fs.String("http_port", "9000", "application http port")
		serviceName = fs.String("service_name", "sloth_service", "service name")
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

	svc := sloth.NewLogMW(logger, sloth.NewService(sloth.NewRepository(), nc.GetConn(), sloth.NewEvents()))

	routes := sloth.NewHTTPServer(svc, logger)

	logger.Log(
		"service name", *serviceName,
		"msg", "HTTP",
		"addr", *httpPort,
		"prom-path", "/metrics")

	logger.Log("err", http.ListenAndServe(":"+*httpPort, routes))

}

func createLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
