variable "subnet_id" {
  description = "The ID of the Subnet"
  type        = string
}

variable "security_group_id" {
  description = "The ID of the Security Group"
  type        = string
}

variable "key_pair_name" {
  description = "the key pair for the ec2"
  type        = string
}