# Event Driven applications with Kinesis and Golang

## Architecture

![Architecture](./architecture/kinesis-lambda.drawio.png)

Each shard can support writes up to 1,000 records per second, up to a maximum data write total of 1 MB per second. Each PutRecords request can support up to 500 records. Each record in the request can be as large as 1 MB, up to a limit of 5 MB for the entire request, including partition keys.
