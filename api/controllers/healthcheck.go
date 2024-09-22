/* =========================================================================
*  File Name: controller/healthcheck.go
*  Description: Controller for the healthcheck route.
*  Author: MagnusChase03
*  =======================================================================*/
package controllers

import (
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Sends simple ok for API healthcheck.
*
*  Arguments:
*      - N/A
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*/
func HealthcheckController() utils.JSONResponse { 
    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    };
}
