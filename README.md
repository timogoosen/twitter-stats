
My data flows this way:

twitter API <----producer--->Dynamodb <- Consumer




```

echo 2
```

I would like to figure out with this project how to:

1. Get data out of Dynamodb.

2. Do something interesting with the data.

3. How to work with a LSI

4. How to use a GSI

5. How to use a dynamodb stream enabled table and act on specific
types of updates in the table.

6. How to cache read queries with Redis and adapting your code to query redis instead.

7. How to invalidate your redis cache by reading from stream.

8. Figure out how to work with Kinesis and write my own producer

9. Figure out how to write my own Kinesis consumer without the KCL


## Dependencies:


1. twitter streaming script:

```
$ pip install tweepy --user


```

If you keep getting import errors for tweepy then you might have to run this:
(On Ubuntu pip is now aliased to pip3)
```

$ sudo pip2 install tweepy --upgrade


```

## Usage:

Export keys as environment variables:

```
export CONSUMER_KEY=INSERTHERE
export CONSUMER_SECRET=INSERTHERE
export ACCESS_TOKEN=INSERTHERE
export ACCESS_TOKEN_SECRET=INSERTHERE
```

Now run twitter stream script and pipe json one document at a time to custom Kinesis producer
like this:

```

$  python twitter2.py | jq -c '.' | go run kinesis-producer.go

```
