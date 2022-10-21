resource "aws_kinesis_stream" "kinesis_data_stream" {
  name             = "${var.kinesis_stream_name}-${var.env}"
  shard_count      = var.kinesis_shard_count
  retention_period = 24
  encryption_type  = "KMS"
  kms_key_id       = "alias/aws/kinesis"
  shard_level_metrics = [
    "IncomingBytes",
    "OutgoingBytes",
  ]
  stream_mode_details {
    stream_mode = "PROVISIONED"
  }
  tags = {
    Name        = "${var.kinesis_stream_name}-${var.env}"
    Environment = var.env
  }
}
