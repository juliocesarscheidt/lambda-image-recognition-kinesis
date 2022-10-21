resource "aws_s3_bucket" "bucket" {
  bucket = "${var.bucket_name}-${var.env}"
  acl    = "private"
  # with versioning enabled
  versioning {
    enabled = true
  }
  lifecycle_rule {
    id      = "old-files"
    enabled = true
    # prefix  = ""
    # move to one zone infrequent access
    transition {
      days          = 30
      storage_class = "ONEZONE_IA"
    }
    # remove older versions of objects
    noncurrent_version_expiration {
      days = 90
    }
  }
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256" # SSE-S3
      }
    }
  }
  tags = {
    Name        = "${var.bucket_name}-${var.env}"
    Environment = var.env
  }
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.bucket.id
  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda_producer.arn
    events              = ["s3:ObjectCreated:*"]
    filter_suffix       = ".png"
  }
  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda_producer.arn
    events              = ["s3:ObjectCreated:*"]
    filter_suffix       = ".jpg"
  }
  depends_on = [
    aws_s3_bucket.bucket,
    aws_lambda_function.lambda_producer,
    aws_lambda_permission.lambda_producer_trigger_from_bucket,
  ]
}

# resource "aws_s3_bucket_policy" "bucket_policy" {
#   bucket = aws_s3_bucket.bucket.id
#   policy = <<EOF
# {
#   "Version": "2012-10-17",
#   "Statement": [
#         {
#             "Sid": "DenyIncorrectEncryptionHeader",
#             "Effect": "Deny",
#             "Principal": "*",
#             "Action": "s3:PutObject",
#             "Resource": "arn:aws:s3:::${var.bucket_name}-${var.env}/*",
#             "Condition": {
#                 "StringNotEquals": {
#                     "s3:x-amz-server-side-encryption": "AES256"
#                 }
#             }
#         },
#         {
#             "Sid": "DenyUnencryptedObjectUploads",
#             "Effect": "Deny",
#             "Principal": "*",
#             "Action": "s3:PutObject",
#             "Resource": "arn:aws:s3:::${var.bucket_name}-${var.env}/*",
#             "Condition": {
#                 "Null": {
#                     "s3:x-amz-server-side-encryption": "true"
#                 }
#             }
#         }
#     ]
# }
# EOF
#   depends_on = [
#     aws_s3_bucket.bucket,
#   ]
# }
