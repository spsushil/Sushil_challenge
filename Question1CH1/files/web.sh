#!/bin/bash

sudo apt update

# Install Apache2
sudo apt install -y apache2
mkdir /etc/apache2/ssl
openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
              -keyout /etc/apache2/ssl/apache.key \
              -out /etc/apache2/ssl/apache.crt \
              -subj "/C=US/ST=CA/L=Los Angeles/O=MyOrg/OU=MyUnit/CN=myserver.example.com"
cat > /etc/apache2/sites-available/default-ssl.conf << EOF
<IfModule mod_ssl.c>
    <VirtualHost _default_:443>
        ServerAdmin webmaster@localhost
        DocumentRoot /var/www/html
        ErrorLog ${APACHE_LOG_DIR}/error.log
        CustomLog ${APACHE_LOG_DIR}/access.log combined
        SSLEngine on
            SSLCertificateFile      /etc/apache2/ssl/apache.crt
            SSLCertificateKeyFile /etc/apache2/ssl/apache.key
            <FilesMatch "\.(cgi|shtml|phtml|php)$">
                SSLOptions +StdEnvVars
                </FilesMatch>
                    <Directory /usr/lib/cgi-bin>
                    SSLOptions +StdEnvVars
                </Directory>
    </VirtualHost>
</IfModule>
EOF

a2enmod ssl
a2ensite default-ssl
systemctl restart apache2

cat > /var/www/html/index.html << EOF
<html>
<head>
<title>Hello World</title>
</head>
<body>
<h1>Hello World!</h1>
</body>
</html>
EOF
