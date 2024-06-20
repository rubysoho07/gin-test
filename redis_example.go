package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

type RedisData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ConnectRedis() {
	host := os.Getenv("REDIS_HOST")

	if host == "" {
		host = "localhost"
	}

	redis_conn_string := fmt.Sprintf("%s:6379", host)

	rdb = redis.NewClient(&redis.Options{
		Addr:     redis_conn_string,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetDataFromRedis(c *gin.Context) {

	key := c.Param("key")

	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Println(err)
	}

	c.String(http.StatusOK, key+" = "+val)
}

func PutDataToRedis(c *gin.Context) {

	var data RedisData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rdb.Set(context.Background(), data.Key, data.Value, 0)

	c.String(http.StatusOK, "Data stored successfully")
}
