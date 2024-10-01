/* =========================================================================
*  File Name: middleware/cors.go
*  Description: The CORS middleware
*  Author: MagnusChase03
*  =======================================================================*/
package middleware

import (
    "net/http"
)

/*
*  Sets CORS Headers on responses.
*
*  Arguments:
*      - next (http.Handler): The next handler to call.
* 
*  Returns:
*      - http.Handler: The main handler with logging attached.
*/
func CorsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "https://localhost:8080");
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS");
        w.Header().Set("Access-Control-Allow-Credentials", "true");
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization");
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusOK);
            return;
        }

        next.ServeHTTP(w, r);
    });
}
