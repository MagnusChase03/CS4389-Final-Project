/* =========================================================================
*  File Name: routes/userRoutes/createUser.go
*  Description: Handler for creating users.
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
    "github.com/MagnusChase03/CS4389-Project/controllers/userControllers"
)

/*
*  Handles the control flow for the create user route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func CreateUserHandler(w http.ResponseWriter, r *http.Request) { 
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
    publicKey := r.FormValue("publicKey");
    if username == "" || password == "" || publicKey == "" {
        fmt.Printf("[ERROR] username, password, or public key empty.\n");
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

    resp, err := userControllers.CreateUserController(username, password, publicKey);
    if err != nil {
        fmt.Fprintf(os.Stderr, "[ERROR] %v\n", err);
    }

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
