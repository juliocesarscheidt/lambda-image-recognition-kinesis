data "archive_file" "lambda_producer_zip" {
  type        = "zip"
  source_file = var.lambda_producer_file_name
  output_path = "${var.lambda_producer_file_name}.zip"
}

resource "aws_lambda_function" "lambda_producer" {
  function_name    = "${var.lambda_producer_name}-${var.env}"
  filename         = "${var.lambda_producer_file_name}.zip"
  handler          = var.lambda_producer_file_name
  source_code_hash = filebase64sha256(data.archive_file.lambda_producer_zip.output_path)
  role             = aws_iam_role.lambda_producer_role.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 30
  environment {
    variables = {
      KINESIS_STREAM_NAME = "${var.kinesis_stream_name}-${var.env}"
    }
  }
  depends_on = [
    aws_iam_role.lambda_producer_role,
    aws_kinesis_stream.kinesis_data_stream,
  ]
  tags = {
    Name        = "${var.lambda_producer_name}-${var.env}"
    Environment = var.env
  }
}

resource "aws_lambda_permission" "lambda_producer_trigger_from_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_producer.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.bucket.arn
  depends_on = [
    aws_lambda_function.lambda_producer,
    aws_s3_bucket.bucket,
  ]
}
