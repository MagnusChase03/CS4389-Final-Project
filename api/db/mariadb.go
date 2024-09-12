/* =========================================================================
*  File Name: db/mariadb.go
*  Description: Handles connected to the MariaDB.
*  Author: MagnusChase03
*  =======================================================================*/
package db;

import (
    "fmt"
    "sync"

    "database/sql"
    _ "github.com/go-sql-driver/mysql"

    "github.com/MagnusChase03/CS4389-Project/utils"
)

// Singleton design pattern
type MariaDB struct {
    Connection *sql.DB;
}

var mariaInstance *MariaDB;
var mariaMutex sync.Mutex;

/*
*  Returns the connection to the MariaDB
*
*  Arguments:
*      - N/A
*
*  Returns:
*      - *MariaDB: The connection to MariaDB..
*/
func GetMariaDB() (*MariaDB, error) {
    if mariaInstance != nil {
        return mariaInstance, nil;
    }

    mutex.Lock();
    if mariaInstance == nil {
        env := utils.GetEnvironment();

        con, err := sql.Open(
            "mysql",
            fmt.Sprintf(
                "%s:%s@(mariadb:3306)/CS4389", 
                env["MARIADB_USER"], 
                env["MARIADB_PASSWORD"],
            ),
        );
        if err != nil {
            mutex.Unlock();
            return nil, fmt.Errorf("[ERROR] Failed to connect to mariadb. %w", err);
        }

        mariaInstance = &MariaDB{con};
    }
    mutex.Unlock();
    return mariaInstance, nil;
}
