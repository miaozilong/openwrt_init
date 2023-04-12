# /etc/config/dhcp
# uci del dhcp.lan.ra_slaac
# uci del dhcp.lan.ra_flags
# /etc/config/network
# uci set network.lan.gateway='192.168.0.1'
# uci add_list network.lan.dns='192.168.0.1'   # TODO  需要重新设置

#配置穿透
# /etc/config/zerotier
uci set zerotier.sample_config.enabled='1'
uci set zerotier.sample_config.nat='0'
uci del zerotier.sample_config.join
uci add_list zerotier.sample_config.join='632ea290858c6736'

# 修改ssh端口
uci set dropbear.@dropbear[0].Port=50022
uci commit dropbear
/etc/init.d/dropbear reload

# 修改web端口
sed -i "s/80/50080/g" /etc/config/nginx
# /etc/init.d/nginx restart

echo '#!/bin/bash

limit_date_file=/etc/limit_date.txt
if [ -f "$limit_date_file" ]; then
    echo "$limit_date_file exist"
else
    echo "2024-04-22 23:45:10" > /etc/limit_date.txt
fi

#一直检测是否过期
for (( ; ; ))
do
	# 30*60 每30分钟检测一次
	sleep 1800
	str_limit=$(cat /etc/limit_date.txt)
	str_now=$(date "+%Y-%m-%d %H:%M:%S")
	date_limit=`date -d "$str_limit" +%s`
	date_now=`date -d "$str_now" +%s`
	if [ $date_limit -gt $date_now ]; then
		echo "没有过期"
	else
		# 过期了直接关机
		halt
	fi
done
' >/bin/check.sh
chmod +x /bin/check.sh
nohup /bin/check.sh >/dev/null 2>&1 &
