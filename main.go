package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var addr = flag.String("addr", "127.0.0.1:8080", "address")
	flag.Parse()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "hello world")
		})
		log.Println("listening on addr", *addr)
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Println("got signal", sig)
	time.Sleep(10 * time.Second)
	log.Println("exiting now")
}
