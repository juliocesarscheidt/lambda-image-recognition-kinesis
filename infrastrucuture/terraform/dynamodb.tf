resource "aws_dynamodb_table" "dynamodb_table" {
  name           = "${var.dynamodb_table_name}_${var.env}"
  hash_key       = "path"
  stream_enabled = false
  billing_mode   = "PROVISIONED"
  read_capacity  = var.dynamodb_rcu
  write_capacity = var.dynamodb_wcu
  attribute {
    name = "path"
    type = "S"
  }
  tags = {
    Name        = "${var.dynamodb_table_name}_${var.env}"
    Environment = var.env
  }
}
