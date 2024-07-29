resource "aws_instance" "web" {
  ami             = "ami-04a81a99f5ec58529" # Ubuntu Server 20.04 LTS (HVM), SSD Volume Type
  instance_type   = "t2.micro"
  subnet_id       = var.subnet_id
  security_groups = [var.security_group_id]
  key_name        = aws_key_pair.deployment_key.key_name


  user_data = <<-EOF
                #!/bin/bash
                apt-get update
                apt-get install -y nginx openssl
                mkdir -p /usr/share/nginx/html
                echo '<html><head><title>Hello World</title></head><body><h1>Hello World!</h1></body></html>' > /usr/share/nginx/html/index.html

                # Configure self-signed SSL certificate
                mkdir -p /etc/nginx/ssl
                openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/selfsigned.key -out /etc/nginx/ssl/selfsigned.crt -subj "/C=US/ST=State/L=City/O=Organization/OU=Org/CN=localhost"
                echo 'server {
                  listen 80;
                  server_name localhost;
                  return 301 https://$host$request_uri;
                }
                server {
                  listen 443 ssl;
                  server_name localhost;
                  ssl_certificate /etc/nginx/ssl/selfsigned.crt;
                  ssl_certificate_key /etc/nginx/ssl/selfsigned.key;
                  location / {
                    root /usr/share/nginx/html;
                    index index.html;
                  }
                }' > /etc/nginx/sites-available/default

                systemctl restart nginx
                EOF
  tags = {
    Name = "WebServerInstance"
  }
}

resource "aws_key_pair" "deployment_key" {
  key_name   = "deployment-key"
  public_key = file("~/.ssh/id_rsa.pub") # Update this path if your key is in a different location
}

resource "aws_eip" "web_ip" {
  instance = aws_instance.web.id
}
