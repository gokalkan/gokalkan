#!/bin/bash

if [ -e /usr/local/share/ca-certificates/extra/ ]; 
then 
	echo "Folder already exists"
else
	mkdir /usr/local/share/ca-certificates/extra
fi

cp -a certs/production/*.crt /usr/local/share/ca-certificates/extra/
cp -a certs/production/*.pem /etc/ssl/certs/
update-ca-certificates

echo "✅ [certs uploaded]"

