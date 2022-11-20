#!/usr/bin/env bash
set -exu

if [ -n "$(git status --porcelain)" ]; then
	echo "Working tree is not clean"
	exit 1
fi

YEAR=$(printf "%d" ${YEAR})
DAY=$(printf "%d" ${DAY})
LPADDAY=$(printf "%02d" ${DAY})

PACKAGE=y${YEAR}d${LPADDAY}
mkdir $PACKAGE
sed -i '' -e "/${PACKAGE}/s/\/\/ //" $(git rev-parse --show-toplevel)/cmd/solve/init.go
cp -r template/* $PACKAGE 
find ./${PACKAGE} -type f -name '*.go' -exec sed -i '' "s/{YEAR}/${YEAR}/g" {} \;
find ./${PACKAGE} -type f -name '*.go' -exec sed -i '' "s/{DAY}/${DAY}/g" {} \;
find ./${PACKAGE} -type f -name '*.go' -exec sed -i '' "s/{LPADDAY}/${LPADDAY}/g" {} \;
git add -A
git commit -m "start $PACKAGE"
