#!/bin/sh

# Request input from the user
read -p "Enter your server_name (example.com , localhost): " server_name

# Store the input in a variable
sn=$server_name

# Request input from the user
read -p "Enter the port for your server: (80, 4000 ...)" server_port

# Store the input in a variable
sp=$server_port

# Create config string
nginx_config="server {
    listen $sp;
    listen [::]:$sp;
    server_name $sn;

    location / {
        root /var/www/html/libr34;
        index index.html;
    }

    location /api {
        proxy_pass http://localhost:8080;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
    }
}"

if [ ! -d "/var/www/html/libr34" ]; then
    mkdir -p /var/www/html/libr34
fi

cp -r ./site/* /var/www/html/libr34

echo -e "Config: \033[32m$nginx_config\033[0m"

if [ ! -d "/etc/nginx/sites-available" ]; then
    mkdir -p /etc/nginx/sites-available
fi

touch /etc/nginx/sites-available/libr34

echo "$nginx_config" > /etc/nginx/sites-available/libr34

if [ ! -d "/etc/nginx/sites-enabled" ]; then
    mkdir -p /etc/nginx/sites-enabled
fi

if [ ! -f "/etc/nginx/sites-enabled/libr34" ]; then
    ln -s /etc/nginx/sites-available/libr34 /etc/nginx/sites-enabled/libr34
fi

