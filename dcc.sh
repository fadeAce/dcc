#!/usr/bin/env bash

FILENAME=$PWD"/data/sock.dat"
OSV=`uname`

while true;
do
$(sed -n '1p' ${FILENAME});
if [[ ${OSV} == "Darwin" ]]; then
sed -i '' '1d' ${FILENAME};
else sed -i '1d' ${FILENAME};
fi
sleep 5;
done