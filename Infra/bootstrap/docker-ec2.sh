#!/bin/bash
sudo yum update -y
sudo yum upgrade -y
sudo yum install -y docker
sudo systemctl start docker
sudo systemctl enable docker