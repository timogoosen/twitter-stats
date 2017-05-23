package main

import (
	//	"encoding/json"
	"fmt"
	"log"
	"sort"
	//	"os"
	//	"strconv"
	"time"
	//"reflect"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

// Retrieve shards

//func GetAllShards(){

func GetFirstShardIteratorForShard(shard_id string, shard_iterator_type string, stream_name string, svc *kinesis.Kinesis) (string, error) {

	params := &kinesis.GetShardIteratorInput{
		ShardId:           aws.String(shard_id),            // Required
		ShardIteratorType: aws.String(shard_iterator_type), // Required
		StreamName:        aws.String(stream_name),         // Required
		Timestamp:         aws.Time(time.Now()),            // good to have
	}
	resp, err := svc.GetShardIterator(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return "", err // cant return nil as type string so just return empty string
	}

	return aws.StringValue(resp.ShardIterator), nil

}

//func GetRecordsAndNextShardIterator( shard_iterator string, svc *kinesis.Kinesis)

func GetShardIDs(streamname string, svc *kinesis.Kinesis) ([]string, error) {

	//func GetShardIDs(streamname string, svc *kinesis.Kinesis) (int, error)

	//	var shards []string

	params := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamname), // Required
		// ExclusiveStartShardId: aws.String("ShardId"),
		Limit: aws.Int64(10000), // Changed this value to 10 now we can see details for atleast 10 shards.
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
/*
	shardIDs := make([]string, len(shards))
	for i, s := range shards {
		shardIDs[i] = aws.StringValue(s.ShardId)
	}
	sort.Strings(shardIDs)
*/
	for i, s := range resp.StreamDescription.Shards {
		shards[i] = aws.StringValue(s.ShardId)  // some of this code I got inspired by twitchscience's kinsumer when dealing with emtpy values in the describe 
// response 
// Some of the code I used for ideas: https://github.com/twitchscience/kinsumer/blob/15fa14b79284d0e66ab2e2960389995eb6a61c8c/leader.go
// This code should be adapted later on to use the pagination for the StreamDescription for when the number of shards
// exceed the limit on the outputof the StreamDescription response as seen in twitchscience's code.

//For some reason we no longer get empty shard ids in the response. Need to test further 
// We could do something to check for nil instead of for "" when looking at the responses, but do this before converting the response to string with aws.StringValue
// this might fix it. This is just a side note for myself for later.
	}
	sort.Strings(shards)
	//return shards, nil;
	return shards, nil



}

func main() {
	// Stream name
	streamname := "kinesis-stream-demo-stream"
	shard_iterator_type := "LATEST"

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

		fmt.Println(shard_ids[v]) // check shard id size that the shard_id has to be minimum field size of 1
				shard_iterator, err := GetFirstShardIteratorForShard(shard_ids[v], shard_iterator_type, streamname, svc)

				// Now get records and nextsharditerator print out types so we can figure out how to write functions
				if err != nil {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())

				}
								
				fmt.Println(shard_iterator) 

	}

	// Now Loop over shard_ids and for each shard_id get a ShardIterator

	// Get shard iterator for each shard id:

}










