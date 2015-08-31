package main

import (
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
)

type Queue struct {
    RedisClient *redis.Pool
}

func (rs *Queue) Init() {
    rs.RedisClient = &redis.Pool{
        MaxIdle:     5,
        MaxActive:   10,
        IdleTimeout: 180 * time.Second,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", "127.0.0.1:6379")
            if err != nil {
                return nil, err
            }
            c.Do("SELECT", 0)
            return c, nil
        },
    }
}

func (rs *Queue) Pop(channel string) (r string, err error) {
    rc := rs.RedisClient.Get()
    defer rc.Close() 
    r, err = redis.String(rc.Do("lpop", channel))
    if err != nil {
        return "", err
    }
    return r, nil
}

func (rs *Queue) Push(channel string, s string) (err error) {
    rc := rs.RedisClient.Get()
    defer rc.Close() 
    _, err = rc.Do("rpush", channel, s)
    if err != nil {
        return err
    }
    return nil
}

func main() {
    q := new(Queue)
    q.Init()
    fmt.Println(q.Push("aa", "123"))
}


