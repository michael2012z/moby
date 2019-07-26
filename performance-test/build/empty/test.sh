#!/bin/bash

START_TIME=$(date)
x=1
while [ $x -le 10 ]
do
  date; docker build -t michael_performance_test:build_empty .;date; docker image rm michael_performance_test:build_empty
  x=$(( $x + 1 ))
done
FINISH_TIME=$(date)
echo ${START_TIME}
echo ${FINISH_TIME}
