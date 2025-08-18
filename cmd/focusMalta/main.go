package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	InfoLog  *slog.Logger
	ErrorLog *slog.Logger
}

func main() {
	// logger handlers
	infoHanlder := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: slog.LevelInfo,
		})

	errorHanlder := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level:     slog.LevelError,
			AddSource: true,
		})

	// loggers.
	infoLog := slog.New(infoHanlder)
	errorLog := slog.New(errorHanlder)

	// application dependencis
	app := &application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	// contain  server configurations.
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: app.routes(),
	}

	// Server Runing massage.
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
