package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AssassinAsh/newsletter/pkg/constants"
	"github.com/AssassinAsh/newsletter/pkg/server"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()

	server.RegisterServer(srv)

	wrappedGrpc := grpcweb.WrapServer(srv,
		grpcweb.WithAllowedRequestHeaders([]string{constants.HeaderOriginValue}),
		grpcweb.WithWebsockets(false))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(constants.HeaderOriginKey, constants.HeaderOriginValue)
		w.Header().Set(constants.HeaderMethodsKey, constants.HeaderMethodsValue)
		w.Header().Set(constants.HeaderNamesKey, constants.HeaderNamesValue)
		w.Header().Set(constants.HeaderExposeKey, constants.HeaderExposeValue)
		if r.Method == constants.RMethodOptions {
			return
		}

		if wrappedGrpc.IsGrpcWebRequest(r) {
			wrappedGrpc.ServeHTTP(w, r)
		} else {
			// Fall back to other servers.
			http.DefaultServeMux.ServeHTTP(w, r)
		}
	})

	port, _ := strconv.Atoi("8080")

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	log.Println("Starting server on port :", port)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("failed starting http server: %v", err)
	}
}
