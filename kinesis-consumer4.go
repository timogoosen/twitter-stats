package main

import (
	//	"encoding/json"
	"fmt"
	"log"
	//	"os"
	//	"strconv"
	//	"time"
	//"reflect"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// Retrieve shards

//func GetAllShards(){

func GetShardIDs(streamname string, svc *kinesis.Kinesis) ([]string, error) {

	//func GetShardIDs(streamname string, svc *kinesis.Kinesis) (int, error)

	//	var shards []string

	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamname), // Required
		// ExclusiveStartShardId: aws.String("ShardId"),
		Limit: aws.Int64(1000), // Changed this value to 10 now we can see details for atleast 10 shards.
		// Need to fix code to be able to handle amount of shards beyond this limit
	}
	resp, err := svc.DescribeStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		//		log.Fatal(err)
		return nil, err
	}

	// Pretty-print the response data.
	//	fmt.Println(resp)

	//Response is obviously some kind of struct that can be accessed to get access to values stored in struct

	status := aws.StringValue(resp.StreamDescription.StreamStatus)
	if status != "ACTIVE" {
		fmt.Errorf("stream %s exists but state '%s' is not 'ACTIVE'", streamname, status)
	}
	fmt.Println(status)
	//shardid := aws.StringValue(resp.StreamDescription.ShardId)
	shards := make([]string, len(resp.StreamDescription.Shards))

	// Looping over map here see example of iterating over a the contents of a map here:
	// https://blog.golang.org/go-maps-in-action
	for _, shard := range resp.StreamDescription.Shards {
		shardid := aws.StringValue(shard.ShardId)
		//fmt.Println(shardid)
		shards = append(shards, shardid)

	}
	//return shards, nil;
	return shards, nil

}

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
	shard_ids, err := GetShardIDs(streamname, svc)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Fatal(err)
	}
	//fmt.Println(shard_ids)
	for v := range shard_ids {

		fmt.Println(shard_ids[v])

	}

	// Now Loop over shard_ids and for each shard_id get a ShardIterator

	// Get shard iterator for each shard id:

}
