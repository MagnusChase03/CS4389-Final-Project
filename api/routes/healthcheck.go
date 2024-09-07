/* =========================================================================
*  File Name: routes/healthcheck.go
*  Description: Handler for the healthcheck route.
*  Author: MagnusChase03
*  =======================================================================*/
package routes

import (
    "net/http"

    "github.com/MagnusChase03/CS4389-Project/utils"
    "github.com/MagnusChase03/CS4389-Project/controllers"
)

/*
*  Handles the control flow for the healthcheck route.
*
*  Arguments:
*      - w (http.ResponseWriter): The object that is used to write a response.
*      - r (*http.Request): The request being made from the client.
* 
*  Returns:
*      - N/A
*/
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) { 
    resp := controllers.HealthcheckController();

    if err := utils.SendResponse(w, resp); err != nil {
        utils.SendInternalServerError(w, err);
    }
}
