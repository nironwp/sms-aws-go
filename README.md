# SMS AWS Go

## Introduction

This is an example Go application for sending SMS messages using AWS SNS.

## Prerequisites

Before running the application, you will need to:

- Create an AWS IAM user with AmazonSNSFullAccess permission.
- Create an AWS access key for the user.
- Update the `.env` file with your AWS access key ID and secret access key.

## Running with Docker Compose

1. Build the Docker image:

   ```shell
   docker build -t sms-aws-go .
   ```

2. Start the Docker container with Docker Compose:
   ```shell
   docker compose up
   ```
