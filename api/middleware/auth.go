/* =========================================================================
*  File Name: middleware/auth.go
*  Description: The authentication middleware
*  Author: MagnusChase03
*  =======================================================================*/
package middleware

import (
    "fmt"
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/models"
    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/session"
)

/*
*  Ensures the user is logged in before proceeding.
*
*  Arguments:
*      - next (http.Handler): The next handler to call.
* 
*  Returns:
*      - http.Handler: The main handler with logging attached.
*/
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("authCookie");
        if err != nil {
            fmt.Printf("[ERROR] Invalid auth cookie. %v\n", err);
            utils.SendUnauthorizedRequest(w);
            return;
        }

        if err = cookie.Valid(); err != nil {
            fmt.Printf("[ERROR] Invalid auth cookie. %v\n", err);
            utils.SendUnauthorizedRequest(w);
            return;
        }

        userID, username, err := session.ParseUserCookie(cookie.Value);
        if err != nil {
            fmt.Printf("[ERROR] Invalid auth cookie. %v\n", err);
            utils.SendUnauthorizedRequest(w);
            return;
        }

        user, err := models.GetUserByID(userID);
        if err != nil {
            fmt.Printf("[ERROR] Invalid auth cookie. %v\n", err);
            utils.SendUnauthorizedRequest(w);
            return;
        }

        if user.Username != username {
            utils.SendUnauthorizedRequest(w);
            return;
        }

        next.ServeHTTP(w, r);
    });
}
