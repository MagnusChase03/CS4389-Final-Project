/* =========================================================================
*  File Name: db/redis.go
*  Description: Handles connected to the Redis DB.
*  Author: MagnusChase03
*  =======================================================================*/
package db;

import (
    "fmt"
    "sync"
    "context"
    "github.com/go-redis/redis/v8"
)

// Singleton design pattern
type RedisDB struct {
    Connection *redis.Client;
    Ctx context.Context;
}

var instance *RedisDB;
var mutex sync.Mutex;

/*
*  Returns the connection to the Redis DB.
*
*  Arguments:
*      - N/A
*
*  Returns:
*      - *RedisDB: The connection to Redis..
*/
func GetRedisDB() (*RedisDB, error) {
    if instance != nil {
        return instance, nil;
    }

    mutex.Lock();
    if instance == nil {
        instance = &RedisDB{
            Connection: redis.NewClient(&redis.Options{
                Addr: "redis:6379",
                Password: "",
                DB: 0,
            }),
            Ctx: context.Background(),
        }; 

        _, err := instance.Connection.Ping(instance.Ctx).Result();
        if err != nil {
            instance = nil;
            mutex.Unlock();
            return nil, fmt.Errorf("[ERROR] Failed to ping to Redis DB. %w", err);
        }
    }
    mutex.Unlock();
    return instance, nil;
}
