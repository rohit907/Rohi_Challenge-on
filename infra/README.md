# Terraform Project for Deploying a Static Web Application on AWS

This project uses Terraform to provision a scalable and secure static web application on AWS. The infrastructure includes an EC2 instance running Nginx with a self-signed SSL certificate to serve a static `index.html` page.

## Prerequisites

- **Terraform:** Install Terraform from [terraform.io](https://www.terraform.io/downloads.html).
- **AWS CLI:** Install the AWS CLI from [aws.amazon.com/cli](https://aws.amazon.com/cli/).
    - Replace <AWS_ACCESS_KEY> with your AWS_ACCESS_KEY in provider.tf
    - Replace <<AWS_ACESS_SECRET_KEY> with yout AWS_ACESS_SECRET_KEY in provider.tf
- **SSH Key Pair:** Ensure you have an existing SSH key pair (`id_rsa` and `id_rsa.pub`) or generate one using `ssh-keygen`.

## Project Structure
```
.
│
├── main.tf
├── modules/
│   ├── vpc/
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   ├── subnet/
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   ├── internet_gateway/
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   ├── route_table/
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   ├── security_group/
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   └── ec2/
│       ├── main.tf
│       ├── outputs.tf
│       └── variables.tf
├── variables.tf
├── outputs.tf
├── test/
│   └── main_test.go
└── README.md
 ```
## Setup Instructions

### 1. Initialize Terraform

Initialize Terraform to download the necessary providers and modules:

```sh
terraform init
```

### 2. Format Configuration Files (Optional)

Format your Terraform configuration files to ensure they are consistent and readable. This step is optional but recommended for maintaining clean code.

```sh
terraform fmt -recursive
```

### 3. Validate the Configuration

Check whether the configuration is syntactically valid and internally consistent. This step ensures that your configuration is correct before you proceed with further actions.

```sh
terraform validate
```



### 4. Plan the Changes

Generate and review an execution plan. This step lets you see what changes Terraform will make before applying them. Review the plan to ensure it meets your expectations.

```sh
terraform plan
```

### 5. Apply Terraform Configuration

Run the following command to apply the Terraform configuration and create the resources:

```sh
terraform apply
```
To skip the confirmation prompt and apply changes automatically, use:

```sh
terraform apply -auto-approve
```

### 6. Obtain the EC2 Instance IP

After Terraform has finished applying, it will output the public IP address of the EC2 instance. Use this IP to access your web server.

https://<instance_public_ip>

Replace <instance_public_ip> with the IP address output by Terraform.


### 7. SSH into the EC2 Instance

If you need to SSH into the EC2 instance for any reason, use the following command:

```sh
ssh -i ~/.ssh/id_rsa ubuntu@<instance_public_ip>
```
Replace <instance_public_ip> with the IP address output by Terraform.


### 8. Verify the Deployment

Navigate to http://<instance_public_ip> in a web browser. You should see the following page:

```sh
<html>
<head>
<title>Hello World</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>
```

#### 9. Clean Up Resources

To destroy the resources created by Terraform and clean up your environment, use:

```sh
terraform destroy
```



#### 10. Output

For reference

![Screenshot 2024-07-29 113515](https://github.com/user-attachments/assets/ce0e05ed-ef05-4001-951e-9074ee0c7632)
![Screenshot 2024-07-29 113651](https://github.com/user-attachments/assets/597010c3-dd89-4abf-866a-4d064f2e5f01)
