#!/bin/bash
limit_date_file=/mofahezi/limit_date.txt
mkdir /tmp/log -p
if [ -f "$limit_date_file" ]; then
  echo "限定日期已存在" >/tmp/log/check.log
else
  cur_sec=$(date '+%s')
  limit_sec=$((cur_sec + 375 * 24 * 60 * 60))
  limit_date=$(date -d @${limit_sec})
  str_limit=$(date -d @$limit_sec "+%Y-%m-%d %H:%M:%S")
  echo "${str_limit}" >${limit_date_file}
  echo "限定日期不存在，写入日期:${str_limit}" >/tmp/log/check.log
fi
#一直检测是否过期
while true; do
  # 30*60 每30分钟检测一次
  sleep 1800
  str_limit=$(cat ${limit_date_file})
  str_now=$(date "+%Y-%m-%d %H:%M:%S")
  date_limit=$(date -d "$str_limit" +%s)
  date_now=$(date -d "$str_now" +%s)
  if [ $date_limit -gt $date_now ]; then
    echo "现在时间:${str_now},过期时间:${str_limit}没有过期" >>/tmp/log/check.log
  else
    echo "现在时间:${str_now},过期时间:${str_limit}过期了直接关机" >>/tmp/log/check.log
    halt
  fi
done
