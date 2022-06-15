#!/usr/bin/env bash

SDK_PATH=$1
ENV=$2 # test или production

if [ -z "$SDK_PATH" ] ; then
    >&2 echo "Ошибка: путь к SDK не указан"
    exit 1
fi

if [ -z "$ENV" ] ; then
    >&2 echo "Ошибка: тип окружения не указан - должно быть test или production"
    exit 1
fi

SCRIPT_PATH="$(dirname "$(realpath "$0")")"
LIB_PATH="$SCRIPT_PATH/../lib/"
SDK_KALKANCRYPT="$SDK_PATH/C/Linux/libs_for_linux/kalkancrypt"
SDK_KALKANCRYPT_SO="$SDK_PATH/C/Linux/C/libkalkancryptwr-64.so"

# удалить содержимое lib и старых сертификатов
rm -f /usr/local/share/ca-certificates/extra/{nca,root}*
rm -f /etc/ssl/certs/{root,nca}*
rm -rf "$LIB_PATH"/*

set -euo pipefail

# разархивировать сертификаты по окружению
unzip -o "$SDK_PATH/C/Linux/ca-certs/$ENV.zip" -d "$SCRIPT_PATH/../ca-$ENV"

# установить доверенные сертификаты
cd "$SCRIPT_PATH/../ca-$ENV/" && bash "install_$ENV.sh" && cd -

# скопировать библиотеку kalkancrypt
[ -d "$LIB_PATH" ] || mkdir -p "$LIB_PATH"
cp -rf "$SDK_KALKANCRYPT" "$LIB_PATH"
cp -f "$SDK_KALKANCRYPT_SO"* "$LIB_PATH"
