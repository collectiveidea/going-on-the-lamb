## Go on Lambda
This project demonstrates how to execute a basic Lambda function written in Go on AWS.

It follows the tutorial from the [Go Support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/) announcement.


## Setup
#### 1. System Dependencies
```sh
$ brew install direnv
$ brew install dep
$ brew upgrade dep
$ brew install awscli
```

#### 2. Configure direnv
The Makefile needs a few environment variables to run. `direnv` helps to load those variables.
```sh
$ cp .envrc{.example,}
$ direnv allow .
```
After that, edit the `.envrc` file to 
Visit [https://github.com/direnv/direnv#setup](https://github.com/direnv/direnv#setup) to configure direnv with your shell.

### 3. Install Go dependencies
We manage our Go dependencies with [dep](https://github.com/golang/dep) for this project. To install the project Go dependencies, run the following from the project's root directory:
```sh
$ dep ensure
```

### 4. Configure AWS
If you've already configured the AWS command line tools skip to [Deploying the Lambda Function]().

  1. Get your AWS credentials from the AWS IAM User Portal by clicking:
  [https://console.aws.amazon.com/iam/home](https://console.aws.amazon.com/iam/home) > Users > [your user] > Security credentials > Access Keys

  2. Use those credentials for `AWS Access Key ID` and `AWS Secret Access Key` (the defaults are fine for the rest) when you run:
  ```sh
  $ aws configure
  ```

## Deploying the Lambda Function
### 1. Build the Lambda Package
```sh
$ make      # Creates build/deployment.zip
```
### 2. Deploy the Lambda Package
```sh
$ make create-function    # Configures the function on AWS
```
### 3. Test the Lambda Function
```sh
$ make test-endpoint
```
