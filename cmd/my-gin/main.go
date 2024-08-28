package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	docs "gin-test/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-test/internal/aws"
	"gin-test/internal/cache"
	"gin-test/internal/database"
	"gin-test/internal/templates"
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
	database.ConnectDB()

	defer database.CloseDb()

	// Connect to Redis
	cache.ConnectRedis()

	defer cache.CloseRedis()

	aws.InitAWS()

	r := gin.Default()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/ping", ping)

	r.GET("/list-buckets", aws.ListBuckets)
	r.GET("/get-temp-credential", aws.GetAssumeRole)

	group := r.Group("/database")
	{
		group.GET("/user/:id", database.GetData)
		group.POST("/user/insert", database.InsertData)
		group.POST("/user/delete/:id", database.DeleteData)
		group.POST("/user/update/:id", database.UpdateData)
	}

	group_redis := r.Group("/redis")
	{
		group_redis.GET("/user/:key", cache.GetDataFromRedis)
		group_redis.POST("/user/insert", cache.PutDataToRedis)
	}

	group_ddb := r.Group("/dynamodb")
	{
		group_ddb.GET("/user/:key", aws.GetItem)
		group_ddb.POST("/user/insert", aws.PutItem)
	}

	group_t := r.Group("/template")
	{
		group_t.POST("/ghaw", templates.GetFileFromTemplate)
	}

	r.Run()
}
