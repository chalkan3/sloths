package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"chalkan.github.com/internal/database"
	"chalkan.github.com/internal/jump"
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

	gorm, err := database.NewPostgresGORM("host=localhost user=postgres password=postgres dbname=jump port=5433 sslmode=disable").Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	jump.NewMigrations(gorm).Migrate()

	svc := jump.NewInstrumentingMiddleware(
		jump.NewMetrics(),
		jump.NewLoggerMW(logger, jump.NewService(jump.NewRepository(gorm))),
	)

	logger.Log(
		"service name", "Jump",
		"msg", "HTTP",
		"addr", "9001",
		"prom-path", "/metrics")

	routes := jump.NewHTTPServer(svc, logger)

	go http.ListenAndServe(":9001", routes)
	jump.SubscriberTransport(svc, nc.GetConn())

}

func createLogger(port string) log.Logger {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stderr)
	logger = log.With(logger, "listen", port, "caller", log.DefaultCaller)
	return logger
}
