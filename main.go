package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error { // web server
		server := http.Server{
			Addr:    "localhost:8081",
			Handler: nil,
		}
		go func() {
			<-ctx.Done()
			fmt.Println("gracefully shutting done web server now")
			_ = server.Shutdown(ctx)
		}()
		return server.ListenAndServe()
	})

	g.Go(func() error { // signal handler
		c := make(chan os.Signal, 2)
		signal.Notify(c, shutdownSignals...)
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				fmt.Println("received shut down signal")
				return errors.New("received shut down signal")
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
