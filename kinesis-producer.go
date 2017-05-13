package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"strconv"
	//"reflect"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {
	// Stream name
	streamname := "kinesis-stream-demo-stream"

	// Stuff to setup Session For Kinesis
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("us-west-2")},
		Profile: "default",
	})

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// process SDK error
			panic(err.Error())
			fmt.Println(awsErr)
		}
	}

	// Open up the session outside the loop so we dont have to open each time
	svc := kinesis.New(sess)

	//
	decoder := json.NewDecoder(os.Stdin)
	decoder.UseNumber()
	jsonstuff := make(map[string]interface{})
	for {
		err := decoder.Decode(&jsonstuff)
		if err != nil {
			//			fmt.Println("caught error")
			//			fmt.Println(err)
		} else {
			fmt.Println("JSON is:")
			//			fmt.Println(jsonstuff)

			result, err := json.Marshal(jsonstuff)
			if err != nil {
				panic(err.Error())
			}
			//fmt.Println(string(result))
			jsonstring := string(result)
			jsonbytes := []byte(jsonstring)
			//fmt.Println(reflect.TypeOf(jsonbytes))
			// uint8 is an alias for byte thats why the output of this is []uint8

			// First convert to string and from string convert to []byte
			timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
			params := &kinesis.PutRecordInput{
				Data:                      []byte(jsonbytes),          // Required
				PartitionKey:              aws.String(timestamp), // Required
				StreamName:                aws.String(streamname),   // Required
			}

			resp, err := svc.PutRecord(params)

			if err != nil {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
				return
			}

			// Pretty-print the response data.
			// sequence number should be in response
			fmt.Println(resp)

		}
	}

}
