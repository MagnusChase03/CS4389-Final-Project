/* =========================================================================
*  File Name: middleware/generic.go
*  Description: Generic definition of middlewares.
*  Author: MagnusChase03
*  =======================================================================*/
package middleware

import (
    "net/http"
)

// Middleware type for ease of use.
type Middleware func(http.Handler) http.Handler

/*
*  Handles the control flow for any handler with middleware attached.
*
*  Arguments:
*      - h (http.Handler): The main route handler
*      - middlewares (Middleware): The list of middlewares to use.
* 
*  Returns:
*      - http.Handler: The main handler with all middlewares attached.
*/
func HandleWithMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
    for _, m := range middlewares {
        h = m(h);
    }

    return h;
}
