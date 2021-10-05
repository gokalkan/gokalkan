#!/bin/bash

set -e

cp -r openssl/kalkancrypt /opt/
cp -f openssl/libkalkancryptwr-64.so /usr/lib/ 

echo "✅ [copy of kalkancrypt libs done]"
