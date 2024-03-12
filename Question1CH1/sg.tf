#security_groups only ssh
resource "aws_security_group" "ssh-sg" {
  name        = "ssh-server"
  description = "Allow ssh traffic"
  vpc_id      = aws_default_vpc.default.id

  # SSH port
  ingress {
    description = "SSH server"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "ssh-sg"
  }
}

#ALB
resource "aws_security_group" "web" {
  name        = "${local.organisation}-web"
  description = "Allow Web inbound traffic"
  vpc_id      = aws_default_vpc.default.id

  ingress {
    description = "HTTP Port"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]

  }

  ingress {
    description = "TLS from Anywhere"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1" 
    cidr_blocks = ["0.0.0.0/0"] 
  }

  tags = {
    Name = "WEB-sg"
  }

  lifecycle {
    create_before_destroy = true
  }
}