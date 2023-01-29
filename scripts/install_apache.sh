#!/bin/sh
echo -e "\033[31mWARNING THIS SCRIPT MAY NOT WORK USE IT WITH THAT IN MIND!\033[0m"
#Request input from the user

read -p "Enter your server_name (example.com , localhost): " server_name
#Store the input in a variable

sn=$server_name
#Request input from the user

read -p "Enter the port for your server: (80, 4000 ...)" server_port
#Store the input in a variable

sp=$server_port
#Create config string

apache_config="<VirtualHost *:$sp>
ServerName $sn
DocumentRoot /var/www/html/libr34

<Directory /var/www/html/libr34>
Options Indexes FollowSymLinks
AllowOverride All
Require all granted
</Directory>

ProxyPass /api/ http://localhost:8080/
ProxyPassReverse /api/ http://$sn/
</VirtualHost>"

if [ ! -d "/var/www/html/libr34" ]; then
mkdir -p /var/www/html/libr34
fi

cp -r ./site/* /var/www/html/libr34

echo -e "Config: \033[32m$apache_config\033[0m"

if [ ! -d "/etc/httpd/sites-available" ]; then
mkdir -p /etc/httpd/sites-available
fi

touch /etc/httpd/sites-available/libr34.conf

echo "$apache_config" > /etc/httpd/sites-available/libr34.conf

if [ ! -d "/etc/httpd/sites-enabled" ]; then
mkdir -p /etc/httpd/sites-enabled
fi

if [ ! -f "/etc/httpd/sites-enabled/libr34.conf" ]; then
ln -s /etc/httpd/sites-available/libr34.conf /etc/httpd/sites-enabled/libr34.conf
fi
echo -e "\033[34mWARNING: Restart the daemon to apply the configuration!\033[0m"