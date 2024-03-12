locals {
  organisation = "sushil"
  vpc_id       = "vpc-0a5f8b6f"
  key_name     = "sushil"
}

resource "aws_default_vpc" "default" {}

resource "aws_default_subnet" "default" {
  availability_zone = "us-east-1a"
}
