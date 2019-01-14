#!/bin/bash

mkdir /usr/local/bin/data
cd /usr/local/bin/data
touch sock.dat

FILENAME="/usr/local/bin/data/sock.dat"
OSV=`uname`

while true;
do
$(sed -n '1p' ${FILENAME});
if [[ ${OSV} == "Darwin" ]]; then
sed -i '' '1d' ${FILENAME};
else sed -i '1d' ${FILENAME};
fi
sleep 3;
done