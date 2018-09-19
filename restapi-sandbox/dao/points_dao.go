package dao

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	. "github.com/dwdii/go-dwdii/restapi-sandbox/models"
)

type PointsDAO struct {
	Server   string
	Database string
}

var db *dynamodb.DynamoDB

func (m *PointsDAO) Connect() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewSharedCredentials("C:\\Users\\Dan\\.aws\\credentials", "default"),
	})

	db = dynamodb.New(sess)

	req := &dynamodb.DescribeTableInput{
		TableName: aws.String("Points"),
	}

	result, err := db.DescribeTable(req)
	if err != nil {
		log.Fatal(err)
	}

	table := result.Table
	log.Printf("%s", table)
}

// Find list of points
func (m *PointsDAO) FindAll() ([]Point, error) {
	var points []Point
	req := &dynamodb.QueryInput{
		TableName: aws.String("Points"),
		KeyConditions: map[string]*dynamodb.Condition{
			"UserId": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("dwdii"),
					},
				},
			},
		},
	}

	var resp1, err1 = db.Query(req)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		err1 = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &points)
	}

	return points, err1
}
