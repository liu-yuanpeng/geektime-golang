package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func myServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)
	server := &http.Server{Addr: ":8080"}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	group.Go(func() error {
		http.HandleFunc("/", myServer)
		fmt.Println("Http Server start ")
		return server.ListenAndServe()
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return server.Shutdown(errCtx)
	})

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-ch:
				//cancel()
				return errors.New("system signal")
			}
		}
		return nil
	})
	//group.Go(func() error {
	//	time.Sleep(time.Second * 100)
	//	fmt.Println("other server")
	//	return errors.New("other error")
	//})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
}
