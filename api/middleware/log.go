/* =========================================================================
*  File Name: middleware/log.go
*  Description: The logging middleware
*  Author: MagnusChase03
*  =======================================================================*/
package middleware

import (
    "fmt"
    "net/http"
)

/*
*  Logs the request information.
*
*  Arguments:
*      - next (http.Handler): The next handler to call.
* 
*  Returns:
*      - http.Handler: The main handler with logging attached.
*/
func LogMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("[LOG] %v - %v.\n", r.Method, r.URL);
        next.ServeHTTP(w, r);
    });
}
