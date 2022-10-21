resource "aws_iam_role" "lambda_consumer_role" {
  name               = "lambda_consumer_role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "lambda_consumer_policy_data_stream" {
  name   = "lambda_consumer_policy_data_stream"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "kinesis:GetRecords",
        "kinesis:GetShardIterator",
        "kinesis:DescribeStream",
        "kinesis:ListShards",
        "kinesis:ListStreams"
      ],
      "Resource": [
        "${aws_kinesis_stream.kinesis_data_stream.arn}",
        "${aws_kinesis_stream.kinesis_data_stream.arn}/*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_consumer_policy_data_stream" {
  role       = aws_iam_role.lambda_consumer_role.name
  policy_arn = aws_iam_policy.lambda_consumer_policy_data_stream.arn
}

resource "aws_iam_policy" "lambda_consumer_policy_cloudwatch" {
  name   = "lambda_consumer_policy_cloudwatch"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_consumer_policy_cloudwatch" {
  role       = aws_iam_role.lambda_consumer_role.name
  policy_arn = aws_iam_policy.lambda_consumer_policy_cloudwatch.arn
}

resource "aws_iam_policy" "lambda_consumer_policy_dynamodb" {
  name   = "lambda_consumer_policy_dynamodb"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "dynamodb:PutItem"
      ],
      "Resource": [
        "${aws_dynamodb_table.dynamodb_table.arn}"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_consumer_policy_dynamodb" {
  role       = aws_iam_role.lambda_consumer_role.name
  policy_arn = aws_iam_policy.lambda_consumer_policy_dynamodb.arn
}
