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
        w.Header().Set("Access-Control-Allow-Origin", "*");
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST");
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization");
        next.ServeHTTP(w, r);
    });
}
