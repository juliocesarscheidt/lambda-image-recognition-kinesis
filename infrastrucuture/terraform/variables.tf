variable "aws_region" {
  type        = string
  description = "AWS region"
  default     = "us-east-1"
}

variable "env" {
  type        = string
  description = "Environment"
  default     = "dev"
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
  default     = "rekognition-bucket-j65hdfa4"
}

variable "dynamodb_table_name" {
  type        = string
  description = "Dynamodb table name"
  default     = "rekognition"
}

variable "lambda_producer_name" {
  type        = string
  description = "Lambda producer name"
  default     = "rekognition-lambda-producer"
}

variable "lambda_producer_file_path" {
  type        = string
  description = "Lambda producer file name"
  default     = "../../lambda-producer/main"
}

variable "lambda_consumer_name" {
  type        = string
  description = "Lambda consumer name"
  default     = "rekognition-lambda-consumer"
}

variable "lambda_consumer_file_path" {
  type        = string
  description = "Lambda consumer file name"
  default     = "../../lambda-consumer/main"
}
