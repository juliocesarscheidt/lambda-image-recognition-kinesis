resource "aws_cloudwatch_log_group" "lambda_producer_log_group" {
  name              = "/aws/lambda/${var.lambda_producer_name}-${var.env}"
  retention_in_days = 1
}

resource "aws_cloudwatch_log_group" "lambda_consumer_log_group" {
  name              = "/aws/lambda/${var.lambda_consumer_name}-${var.env}"
  retention_in_days = 1
}
