#!/usr/bin/env bash
set -exu

YEAR=${YEAR}
DAY=$(printf "%02d" ${DAY})

PACKAGE=y${YEAR}d${DAY}
mkdir $PACKAGE
sed -i '' -e "/${PACKAGE}/s/\/\/ //" $(git rev-parse --show-toplevel)/cmd/solve/init.go
cp -r template/* $PACKAGE 
find ./${PACKAGE} -type f -name '*.go' -exec sed -i '' "s/{YEAR}/${YEAR}/g" {} \;
find ./${PACKAGE} -type f -name '*.go' -exec sed -i '' "s/{DAY}/${DAY}/g" {} \;
