package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var cfg aws.Config
var s3c *s3.Client

func InitAWS() {
	var err error

	cfg, err = config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Println(err)
	}
}

// ListBuckets godoc
// @Summary      List S3 buckets
// @Description  S3 Bucket의 리스트를 가져옵니다.
// @Tags         aws
// @Produce      plain
// @Success      200  {string}  string "bucket names"
// @Router       /list-buckets [get]
func ListBuckets(ctx *gin.Context) {
	if s3c == nil {
		s3c = s3.NewFromConfig(cfg)
	}

	output, err := s3c.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

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
}
