/* =========================================================================
*  File Name: utils/general.go
*  Description: A file containing common functions for all modules.
*  Author: MagnusChase03
*  =======================================================================*/
package utils

import (
    "os"
    "strings"
)


/*
*  Retuns a map of environment variables
*
*  Arguments:
*      - N/A
* 
*  Returns:
*      - map[string]string: The mapping of environment variables
*/
func GetEnvironment() map[string]string {
    result := make(map[string]string);
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=");
        result[pair[0]] = pair[1];
    }
    return result;
}
