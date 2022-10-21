data "archive_file" "lambda_consumer_zip" {
  type        = "zip"
  source_file = var.lambda_consumer_file_name
  output_path = "${var.lambda_consumer_file_name}.zip"
}

resource "aws_lambda_function" "lambda_consumer" {
  function_name    = "${var.lambda_consumer_name}-${var.env}"
  filename         = "${var.lambda_consumer_file_name}.zip"
  handler          = var.lambda_consumer_file_name
  source_code_hash = filebase64sha256(data.archive_file.lambda_consumer_zip.output_path)
  role             = aws_iam_role.lambda_consumer_role.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 30
  environment {
    variables = {
      TABLE_NAME = "${var.dynamodb_table_name}_${var.env}"
    }
  }
  depends_on = [
    aws_iam_role.lambda_consumer_role,
    aws_dynamodb_table.dynamodb_table,
  ]
  tags = {
    Name        = "${var.lambda_consumer_name}-${var.env}"
    Environment = var.env
  }
}

resource "aws_lambda_event_source_mapping" "lambda_consumer_trigger_from_kinesis" {
  event_source_arn  = aws_kinesis_stream.kinesis_data_stream.arn
  function_name     = aws_lambda_function.lambda_consumer.arn
  starting_position = "LATEST"
  depends_on = [
    aws_kinesis_stream.kinesis_data_stream,
    aws_lambda_function.lambda_consumer,
  ]
}
