package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gin-gonic/gin"
)

var ddbc *dynamodb.Client

type GoniTest struct {
	MyKey   string `dynamodbav:"mykey" json:"mykey"`
	Comment string `dynamodbav:"comment" json:"comment"`
}

func InitDDBClient() {
	if ddbc == nil {
		ddbc = dynamodb.NewFromConfig(cfg)
	}
}

func GetItem(ctx *gin.Context) {
	if ddbc == nil {
		InitDDBClient()
	}

	var i GoniTest

	key := ctx.Param("key")

	mykey, err := attributevalue.Marshal(key)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := ddbc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("goni-test"),
		Key: map[string]types.AttributeValue{
			"mykey": mykey,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = attributevalue.UnmarshalMap(response.Item, &i)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, i)
}

func PutItem(ctx *gin.Context) {
	if ddbc == nil {
		InitDDBClient()
	}

	var i GoniTest
	err := ctx.ShouldBindJSON(&i)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := attributevalue.MarshalMap(&i)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err = ddbc.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("goni-test"),
		Item:      item,
	})

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.String(http.StatusOK, "Inserted data with key: %s", i.MyKey)
}
