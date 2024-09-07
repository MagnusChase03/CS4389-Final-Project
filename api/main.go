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
)

func main() {
    // Assign routes
    http.HandleFunc("/healthcheck", routes.HealthcheckHandler);

    // Start HTTPS server on port 8080
    fmt.Printf("[LOG] Starting API server on 0.0.0.0:8080.\n");
    if err := http.ListenAndServeTLS("0.0.0.0:8080", "/certs/server.crt", "/certs/server.key", nil); err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] Failed to start API server. %v\n", err);
    }
}
