package main

import (
  "log"
  "os"

  "github.com/hashicorp/terraform/builtin/provisioners/file"
  "github.com/johandry/terranova"
  "github.com/terraform-providers/terraform-provider-aws/aws"

  "github.com/hashicorp/terraform"
)

var code string

const stateFilename = "aws-ec2-ubuntu.tfstate"

func main() {
  count := 1

  platform, err := terranova.NewPlatform(code).
    AddProvider("aws", aws.Provider()).
    AddProvisioner("file", file.Provisioner()).
    Var("count", count).
    ReadStateFromFile(stateFilename)

  if err != nil {
    if os.IsNotExist(err) {
      log.Printf("[DEBUG] state file %s does not exists", stateFilename)
    } else {
      log.Fatalf("Fail to load the initial state of the platform from file %s. %s", stateFilename, err)
    }
  }

  terminate := (count == 0)
  if err := platform.Apply(terminate); err != nil {
    log.Fatalf("Fail to apply the changes to the platform. %s", err)
  }

  if _, err := platform.WriteStateFile(stateFilename); err != nil {
    log.Fatalf("Fail to save the final state of the platform to file %s. %s", stateFilename, err)
  }
}

func init() {
  code = `
  variable "count"            { default = 2 }
  variable "public_key_file"  { default = "~/.ssh/id_rsa.pub" }
  variable "private_key_file" { default = "~/.ssh/id_rsa" }
  locals {
    public_key    = "${file(pathexpand(var.public_key_file))}"
    private_key   = "${file(pathexpand(var.private_key_file))}"
  }
  provider "aws" {
    region        = "us-west-2"
  }
  resource "aws_instance" "server" {
    instance_type = "t2.micro"
    ami           = "ami-6e1a0117"
    count         = "${var.count}"
    key_name      = "server_key"

    provisioner "file" {
      content     = "ami used: ${self.ami}"
      destination = "/tmp/file.log"

      connection {
        user        = "ubuntu"
        private_key = "${local.private_key}"
      }
    }
  }
  resource "aws_key_pair" "keypair" {
    key_name    = "server_key"
    public_key  = "${local.public_key}"
  }
`
}
