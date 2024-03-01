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
)

func main() {
	// Connect to DB
	ConnectDB()

	defer db.Close()

	// Load AWS credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// AWS Client
	client := s3.NewFromConfig(cfg)

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/list-buckets", func(ctx *gin.Context) {

		// ListBuckets
		output, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

		if err != nil {
			log.Fatal(err)
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

	r.Run()
}
