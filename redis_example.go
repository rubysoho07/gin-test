package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

type RedisData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ConnectRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetDataFromRedis(c *gin.Context) {

	key := c.Param("key")

	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Fatal(err)
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
