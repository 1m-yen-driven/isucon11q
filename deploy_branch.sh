#!/bin/bash

app=()

while getopts "t:r:" optKey; do
 	case $optKey in
	  t)
      echo "--target = $OPTARG"
			app+=("$OPTARG")
      ;;
		r)
			repo=$OPTARG
	esac
done
echo ${app[@]}
echo ${repo}
for a in "${app[@]}"; do
  ssh -o ClearAllForwardings=yes -tt $a /home/isucon/logs/parse.sh skip
  ssh -o ClearAllForwardings=yes $a sudo systemctl stop mysql isucondition.go.service nginx
  ssh -o ClearAllForwardings=yes -tt $a "cd /home/isucon/webapp && git fetch && git checkout ${repo} && cd ./go && make -B" 
  ssh -o ClearAllForwardings=yes $a sudo systemctl start mysql isucondition.go.service nginx
done
