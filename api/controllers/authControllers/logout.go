/* =========================================================================
*  File Name: controller/logout.go
*  Description: Controller for the logout route.
*  Author: MagnusChase03
*  =======================================================================*/
package authControllers

import (
    "github.com/MagnusChase03/CS4389-Project/utils"
)

/*
*  Sends simple ok for logout..
*
*  Arguments:
*      - N/A
* 
*  Returns:
*      - utils.JSONResponse: The response to be made to the client.
*/
func LogoutController() utils.JSONResponse { 
    return utils.JSONResponse{
        StatusCode: 200,
        Data: "Ok",
    };
}
