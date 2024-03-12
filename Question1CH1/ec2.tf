data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "myInstance_sushil" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = "t3a.medium"
  key_name                    = aws_key_pair.deployer.key_name
  user_data_replace_on_change = true
  user_data_base64  = filebase64("files/web.sh")
  subnet_id = aws_default_subnet.default.id
  disable_api_termination = true
  root_block_device {
    delete_on_termination = true
    volume_size           = 15
    volume_type           = "gp2"
  }
  
  
  vpc_security_group_ids = [aws_security_group.ssh-sg.id,aws_security_group.web.id]
  tags = {
    Name = "sushil-newec2"
  }
}

resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = file("C:\\Users\\sushi\\.ssh\\id_rsa.pub")
}

output "ec2_ip" {
  value = aws_instance.myInstance_sushil.public_ip
}


