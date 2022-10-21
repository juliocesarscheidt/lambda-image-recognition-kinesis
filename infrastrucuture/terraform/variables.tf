variable "aws_region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "env" {
  type        = string
  description = "Environment"
  default     = "development"
}

variable "kinesis_stream_name" {
  type        = string
  description = "Kinesis stream name"
  default     = "rekognition-stream"
}

variable "kinesis_shard_count" {
  type        = number
  description = "Kinesis shard count"
  default     = 1
}

variable "bucket_name" {
  type        = string
  description = "Bucket name"
  default     = "rekognition-bucket"
}

variable "dynamodb_table_name" {
  type        = string
  description = "Dynamodb table name"
  default     = "rekognition"
}

variable "dynamodb_rcu" {
  type        = number
  description = "Dynamodb RCU"
  default     = 1
}

variable "dynamodb_wcu" {
  type        = number
  description = "Dynamodb WCU"
  default     = 1
}

variable "lambda_producer_name" {
  type        = string
  description = "Lambda producer name"
  default     = "rekognition-lambda-producer"
}

variable "lambda_producer_file_name" {
  type        = string
  description = "Lambda producer file name"
  default     = "producer"
}

variable "lambda_consumer_name" {
  type        = string
  description = "Lambda consumer name"
  default     = "rekognition-lambda-consumer"
}

variable "lambda_consumer_file_name" {
  type        = string
  description = "Lambda consumer file name"
  default     = "consumer"
}
