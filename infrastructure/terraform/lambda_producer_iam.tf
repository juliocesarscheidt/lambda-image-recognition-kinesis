resource "aws_iam_role" "lambda_producer_role" {
  name               = "lambda_producer_role"
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

resource "aws_iam_policy" "lambda_producer_policy_bucket" {
  name   = "lambda_producer_policy_bucket"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:Get*"
      ],
      "Resource": [
        "${aws_s3_bucket.bucket.arn}/*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_producer_policy_bucket" {
  role       = aws_iam_role.lambda_producer_role.name
  policy_arn = aws_iam_policy.lambda_producer_policy_bucket.arn
}

resource "aws_iam_policy" "lambda_producer_policy_data_stream" {
  name   = "lambda_producer_policy_data_stream"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "kinesis:PutRecord",
        "kinesis:PutRecords"
      ],
      "Resource": [
        "${aws_kinesis_stream.kinesis_data_stream.arn}"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_producer_policy_data_stream" {
  role       = aws_iam_role.lambda_producer_role.name
  policy_arn = aws_iam_policy.lambda_producer_policy_data_stream.arn
}

resource "aws_iam_policy" "lambda_producer_policy_cloudwatch" {
  name   = "lambda_producer_policy_cloudwatch"
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

resource "aws_iam_role_policy_attachment" "attach_role_lambda_producer_policy_cloudwatch" {
  role       = aws_iam_role.lambda_producer_role.name
  policy_arn = aws_iam_policy.lambda_producer_policy_cloudwatch.arn
}

resource "aws_iam_policy" "lambda_producer_policy_rekognition" {
  name   = "lambda_producer_policy_rekognition"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "rekognition:DetectText"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "attach_role_lambda_producer_policy_rekognition" {
  role       = aws_iam_role.lambda_producer_role.name
  policy_arn = aws_iam_policy.lambda_producer_policy_rekognition.arn
}
