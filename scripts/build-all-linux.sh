#!/bin/sh
startTime=`date +%s`
pjPath=$(pwd)
FILE=Makefile
if test -f "$FILE"; then
  echo "$FILE exist"
else
  echo "$FILE not exist, please use './scripts/build.sh adm'"
  exit 0
fi

cd app/cmd
apps=$(ls -d */)
cd ${pjPath}
echo "start build apps"
for i in ${apps}; do
  make build-app-specify-linux name=${i%%/}
done
endTime=`date +%s`

echo "finished in "$[ $endTime - $startTime ]" seconds"
