package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	InfoLog  *slog.Logger
	ErrorLog *slog.Logger
}

func main() {

	// addr flag for define the PORT address
	addr := flag.String("addr", "localhost:8080", "HTTP network address")

	// Parsing all the flags.
	flag.Parse()

	// Handle info logs
	infoHanlder := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		})

	// Handle error logs
	errorHanlder := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     slog.LevelError,
			AddSource: true,
		})

	// loggers.
	infoLog := slog.New(infoHanlder)
	errorLog := slog.New(errorHanlder)

	// application dependencies
	app := &application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// Containing server configurations.
	server := http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	// Server Runing log massage.
	infoLog.Info(
		"Server Running",
		slog.String("PORT", server.Addr),
	)

	err := server.ListenAndServe()
	if err != nil {
		errorLog.Error(err.Error())
		os.Exit(1)
	}
}
