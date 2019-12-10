provider "aws" {
  region = "${var.aws_region}"
  access_key = "${var.AWSAccessKeyId}"
  secret_key = "${var.AWSSecretKey}"
}
data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["calculadoraGo_devops_tema13"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["self"]
}
resource "aws_launch_configuration" "as_conf" {
  name_prefix   = "lc_calculadoraGo_DevOps"
  image_id      = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "bar" {
  name                 = "lc_calculadoraGo_DevOps"
  launch_configuration = "${aws_launch_configuration.as_conf.name}"
  availability_zones = ["${var.aws_zone}"]
  min_size             = 1
  max_size             = 2
  
  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_security_group" "allow_tls" {
  name        = "allow_all"
  description = "Allow TLS inbound traffic"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["200.146.254.130/32"]
  }

  ingress {
    from_port   = 9000
    to_port     = 9000
    protocol    = "tcp"
    cidr_blocks = ["200.146.254.130/32"]
  }

  tags = {
    Name = "allow_all"
  }
}

resource "aws_instance" "web" {
  ami           = "${data.aws_ami.ubuntu.id}"
  instance_type = "t2.micro"
  security_groups = ["${aws_security_group.allow_tls.name}"]

  tags = {
    Name = "sg_calculadora"
  }
}

resource "aws_elb" "mod" {
  name               = "calculadora-terraform-elb"
  availability_zones = ["ca-central-1a", "ca-central-1b"]

  listener {
    instance_port     = 9000
    instance_protocol = "http"
    lb_port           = 80
    lb_protocol       = "http"
  }

  instances = ["${aws_instance.web.id}"]

}

