#!/bin/sh


function print_help() {
	echo "Usages: "
	echo $0 "/path/to *.crt file"
	exit 0
}

if [ $# != 1 ]; then
	print_help
fi

for f in $1/*.crt
do
	openssl x509 -in $f -noout -dates
done
