package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/fernandoocampo/paymentgw/pkg/adapter/web"
	"github.com/fernandoocampo/paymentgw/pkg/application"
)

func main() {
	var httpAddr = flag.String("http", ":8287", "http listen address")
	flag.Parse()
	ctx := context.Background()
	srv := application.NewBasicPaymentProcessor()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	endpoints := web.Endpoints{
		PaymentProcessorEndpoint: web.MakePaymentProcessorEndpoint(srv),
	}

	go func() {
		log.Println("http:", *httpAddr)
		handler := web.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
