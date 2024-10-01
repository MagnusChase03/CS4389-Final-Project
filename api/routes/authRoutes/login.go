/* =========================================================================
*  File Name: routes/login.go
*  Description: Handler for logging in users.
*  Author: MagnusChase03
*  =======================================================================*/
package authRoutes

import (
    "os"
    "fmt"
    "net/http"
    "crypto/sha256"
    "encoding/hex"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers/authControllers"
    "github.com/MagnusChase03/CS4389-Project/session"
)

/*
*  Handles the control flow for the login route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func LoginHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        utils.SendBadRequest(w);
        return;
    }

    err := r.ParseForm();
    if err != nil {
        fmt.Printf("[ERROR] Failed to parse form.\n");
        utils.SendBadRequest(w);
        return;
    }

    username := r.FormValue("username");
    password := r.FormValue("password");
    if username == "" || password == "" {
        fmt.Printf("[ERROR] username of password empty.\n");
        utils.SendBadRequest(w);
        return;
    }

    hasher := sha256.New();
    _, err = hasher.Write([]byte(password));
    if err != nil {
        utils.SendInternalServerError(w, err);
        return;
    }
    password = hex.EncodeToString(hasher.Sum(nil));

    resp, user, err := authControllers.LoginController(username, password);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    } else {
        cookie, err := session.CreateUserCookie(user);
        if err != nil {
            utils.SendInternalServerError(w, err);
            return;
        }
        http.SetCookie(w, &cookie);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
