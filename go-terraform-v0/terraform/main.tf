resource "aws_instance" "test_instance" {
  ami           = "ami-0d1cd67c26f5fca19"
  instance_type = "t2.micro"

  tags = {
    "Env"      = "Private"
    "Location" = "Secret"
  }
}