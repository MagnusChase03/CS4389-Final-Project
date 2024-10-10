/* =========================================================================
*  File Name: routes/userRoutes/updateUser.go
*  Description: Handler for updating user information.
*  Author: MagnusChase03
*  =======================================================================*/
package userRoutes

import (
    "os"
    "fmt"
    "net/http"
    "crypto/sha256"
    "encoding/hex"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/session"
    "github.com/MagnusChase03/CS4389-Project/controllers/userControllers"
)

/*
*  Handles the control flow for the update user route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        utils.SendBadRequest(w);
        return;
    }

    cookie, err := r.Cookie("authCookie");
    if err != nil {
        utils.SendUnauthorizedRequest(w);
        return;
    }

    userID, _, err := session.ParseUserCookie(cookie.Value);
    if err != nil {
        utils.SendUnauthorizedRequest(w);
        return;
    }

    err = r.ParseForm();
    if err != nil {
        fmt.Printf("[ERROR] Failed to parse form.\n");
        utils.SendBadRequest(w);
        return;
    }

    password := r.FormValue("password");
    publicKey := r.FormValue("publicKey");
    if password == "" && publicKey == "" {
        fmt.Printf("[ERROR] Password or public key empty.\n");
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

    resp, err := userControllers.UpdateUserController(userID, password, publicKey);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
