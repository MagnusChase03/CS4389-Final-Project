/* =========================================================================
*  File Name: main.go
*  Description: Starts the HTTP API and hands off requests.
*  Author: MagnusChase03
*  =======================================================================*/
package main

import (
    "os"
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/routes"
    "github.com/MagnusChase03/CS4389-Project/middleware"
)

func main() {
    mux := http.NewServeMux();

    // Assign routes
    mux.Handle("/healthcheck", middleware.HandleWithMiddleware(
        http.HandlerFunc(routes.HealthcheckHandler),
        middleware.LogMiddleware,
    ));

    // Start HTTPS server on port 8080
    fmt.Printf("[LOG] Starting API server on 0.0.0.0:8080.\n");
    //if err := http.ListenAndServeTLS("0.0.0.0:8080", "/certs/server.crt", "/certs/server.key", mux); err != nil {
    if err := http.ListenAndServeTLS("0.0.0.0:8080", "../certs/server.crt", "../certs/server.key", mux); err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] Failed to start API server. %v\n", err);
    }
}
