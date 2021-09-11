#! /bin/bash
echo 'usage: ./parse-all.sh murata-7272'
echo '  or : ./parse-all.sh skip'
echo '  or : ./parse-all.sh '

path=`date +"%I/%M-$1"`

for app in app1 app2 app3; do
  if [[ $1 == "skip" ]]; then
    ssh $app -tt "bash logs/parse.sh skip"
  else
    ssh $app -tt "bash logs/parse.sh $path"
  fi
done

# if [[ $1 != "skip" ]]; then
#   # 必要ならSlackに通知
#   curl -X POST --data-urlencode "payload={\"text\": \"$path のログが更新されました。https://sagisawa.0am.jp/hyaku/logs/$path\"}" https://hooks.slack.com/services/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
# fi
