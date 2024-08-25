package aws

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var cfg aws.Config
var s3c *s3.Client
var stsc *sts.Client

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

func GetAccountId() string {
	if stsc == nil {
		stsc = sts.NewFromConfig(cfg)
	}

	output, err := stsc.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})

	if err != nil {
		log.Println(err)
		return ""
	}

	return aws.ToString(output.Account)
}

func GetAssumeRole(ctx *gin.Context) {

	var result bytes.Buffer
	var role_arn string = fmt.Sprintf("arn:aws:iam::%s:role/gin-test-assume-role", GetAccountId())

	if stsc == nil {
		stsc = sts.NewFromConfig(cfg)
	}

	output, err := stsc.AssumeRole(context.TODO(), &sts.AssumeRoleInput{
		RoleArn:         aws.String(role_arn),
		RoleSessionName: aws.String("gin-test"),
	})

	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	// For ~/.aws/credentials
	result.WriteString(fmt.Sprintln("[profile gin-test]"))
	result.WriteString(fmt.Sprintf("aws_access_key_id = %s\n", *output.Credentials.AccessKeyId))
	result.WriteString(fmt.Sprintf("aws_secret_access_key = %s\n", *output.Credentials.SecretAccessKey))
	result.WriteString(fmt.Sprintf("aws_session_token = %s\n", *output.Credentials.SessionToken))

	// For ~/.aws/config
	result.WriteString(fmt.Sprintln("[profile gin-test]"))
	result.WriteString(fmt.Sprintln("region = ap-northeast-2"))
	result.WriteString(fmt.Sprintln("source_profile = default"))
	result.WriteString(fmt.Sprintf("role_arn = %s\n", role_arn))

	ctx.String(http.StatusOK, result.String())
}
