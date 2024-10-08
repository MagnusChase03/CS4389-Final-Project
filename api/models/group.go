/* =========================================================================
*  File Name: models/group.go
*  Description: Handles database interactions related to groups.
*  Author: Matthew-Basinger
*  =======================================================================*/
package models;

import (
    "fmt"
    "github.com/MagnusChase03/CS4389-Project/db"
)

type Group struct {
    GroupID int
    CreatorID int
    GroupName string
};

/*
*  Returns a group with given GroupID
*
*  Arguments
*      - id (int): The GroupID to find
*
*  Returns: 
*      - Group: the group information
*      - error: An error if any occured
*/

func GetGroupByID(id int) (Group, error) {
    var group Group;
    instance, err := db.GetMariaDB();
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to get mariadb instance. %w", err);
    }

    query, err := instance.Connection.Prepare("SELECT * FROM Groups WHERE GroupID = ?")
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to get parse SQL query. %w", err);
    }
    defer query.Close();

    err = query.QueryRow(id).Scan(
        &group.GroupID, 
        &group.CreatorID, 
        &group.GroupName, 
    );
    if err != nil {
        return group, fmt.Errorf("[ERROR] Failed to find group with GroupID %d. %w", id, err);
    }

    return group, nil;
}

/*
*
*
*
*/

