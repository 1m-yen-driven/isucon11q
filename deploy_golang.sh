#!/bin/bash

cd go
make cross-compile

for app in "$@"; do
  # ssh -o ClearAllForwardings=yes -tt $app /home/isucon/logs/parse.sh skip
  ssh -o ClearAllForwardings=yes $app sudo systemctl stop mysql isucondition.go.service nginx
  scp isucondition $app:/home/isucon/webapp/go/isucondition
  ssh -o ClearAllForwardings=yes $app sudo systemctl start mysql isucondition.go.service nginx
done
