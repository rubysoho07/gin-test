package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
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

	// Load AWS credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println(err)
	}
	// AWS Client
	client := s3.NewFromConfig(cfg)

	r := gin.Default()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/ping", ping)

	r.GET("/list-buckets", func(ctx *gin.Context) {

		// ListBuckets
		output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

		if err != nil {
			log.Println(err)
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		var result []string

		for _, element := range output.Buckets {
			result = append(result, aws.ToString(element.Name))
		}

		ctx.String(http.StatusOK, strings.Join(result, "\n"))
	})

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
