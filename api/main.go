/* =========================================================================
*  File Name: main.go
*  Description: Starts the HTTP API and hands off requests.
*  Author: MagnusChase03, Matthew
*  =======================================================================*/
package main

import (
    "os"
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/routes"
    "github.com/MagnusChase03/CS4389-Project/routes/authRoutes"
    "github.com/MagnusChase03/CS4389-Project/middleware"
    "github.com/MagnusChase03/CS4389-Project/db"
)

func main() {
    // Connect to databases
    _, err := db.GetRedisDB();
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] Failed to connect to Redis DB. %v\n", err);
        return;
    }

    _, err = db.GetMariaDB();
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] Failed to connect to MariaDB. %v\n", err);
        return;
    }

    mux := http.NewServeMux();

    // Assign routes
    mux.Handle("/healthcheck", middleware.HandleWithMiddleware(
        http.HandlerFunc(routes.HealthcheckHandler),
        middleware.CorsMiddleware,
        middleware.LogMiddleware,
    ));

    mux.Handle("/login", middleware.HandleWithMiddleware(
        http.HandlerFunc(authRoutes.LoginHandler),
        middleware.CorsMiddleware,
        middleware.LogMiddleware,
    ));

    mux.Handle("/logout", middleware.HandleWithMiddleware(
        http.HandlerFunc(authRoutes.LogoutHandler),
        middleware.AuthMiddleware,
        middleware.CorsMiddleware,
        middleware.LogMiddleware,
    ));

    // Start HTTPS server on port 8080
    fmt.Printf("[LOG] Starting API server on 0.0.0.0:8080.\n");
    if err := http.ListenAndServeTLS("0.0.0.0:8080", "/certs/server.crt", "/certs/server.key", mux); err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] Failed to start API server. %v\n", err);
    }
}
