package main


import (
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/dynamodb"
"fmt"

)

// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// Might use later

func main() {

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("eu-west-1")},
		Profile: "default",
	})

	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := dynamodb.New(sess)

	// DynamoDB Stuff ends here

	/*
	   resp, err := svc.BatchGetItem(params)

	   if err != nil {

	       fmt.Println(err.Error())
	       return
	   }

	*/
  //

	// Primary key
	// partition key  = ID (String)

	params := &dynamodb.QueryInput{
		TableName:              aws.String("twitter3"),
		IndexName:              aws.String("ID-TweetLang-index"),
		KeyConditionExpression: aws.String("ID = :pkey"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pkey": {
				S: aws.String("819435712441450497"),
			},
		},
	}

	resp, err := svc.Query(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)

}
