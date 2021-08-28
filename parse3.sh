#!/bin/bash

for app in "$@"; do
  ssh -o ClearAllForwardings=yes -tt $app /home/isucon/logs/parse.sh
done
