package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	docs "gin-test/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// PingExample godoc
// @Summary      Show an ping
// @Description  get ping
// @Tags         ping
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /ping [get]
func ping(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	// Connect to DB
	ConnectDB()

	defer db.Close()

	// Connect to Redis
	ConnectRedis()

	defer rdb.Close()

	InitAWS()

	r := gin.Default()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/ping", ping)

	r.GET("/list-buckets", ListBuckets)

	group := r.Group("/database")
	{
		group.GET("/user/:id", GetData)
		group.POST("/user/insert", InsertData)
		group.POST("/user/delete/:id", DeleteData)
		group.POST("/user/update/:id", UpdateData)
	}

	group_redis := r.Group("/redis")
	{
		group_redis.GET("/user/:key", GetDataFromRedis)
		group_redis.POST("/user/insert", PutDataToRedis)
	}

	r.Run()
}
