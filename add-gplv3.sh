#!/bin/bash

for i in $(git ls-files | grep "\.go$" ); do
  echo "Adding GPL3 header to $i"
  cat license-short.txt | sed "s/{open}/\/*/g" | sed "s/{close}/*\//g" | cat - $i > /tmp/temp && mv /tmp/temp $i
done
