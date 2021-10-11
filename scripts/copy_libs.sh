#!/bin/bash

set -e

cp -r lib/kalkancrypt /opt/
cp -f lib/libkalkancryptwr-64.so /usr/lib/ 

echo "✅ [copy of kalkancrypt libs done]"
