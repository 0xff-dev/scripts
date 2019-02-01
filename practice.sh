#!/bin/bash

echo "`whoami`" "now path is `pwd`"

#根据当前的时候判断和生日的月份如果
# 目标日期 1997-06-06

aim_m=6
aim_d=6
now_m=`date +%m`
now_d=`date +%d`
now_year=`date +%Y`

echo `date -j -f %m-%d $now_m-$now_d +%s`

if [ $aim_m == $now_m ]; then
    if [ $aim_d == $now_d ]; then
        echo "你正在过生日"
    elif [ $aim_d > $now_d ]; then
        res=$(( $aim_d-$now_d ))
        echo "差" $res "天"
    else
        res=$(($now_d-$aim_d+365))
        echo "差" $res " 天"
    fi
elif [ $aim_d > $now_m ]; then
    #我们需要
    aim=`date -j -f %m-%d 06-06 +%s`
    now=`date -j -f %m-%d $now_m-$now_d +%s`
    echo "差" $(( ($aim-$now)/60/60/24 )) "天"
else
    # 生日已经度多, 准备下一个
    aim=`date -j -f %Y-%m-%d $((($now_year+1)))-$aim_m-$aim_d +%s`
    now=`date -j -f %Y-%m-%d $now_year-$now_m-$now_d +%s`
    echo "差" $(( ($aim-$now)/60/60/24 )) "天"
fi
# date -j -f %Y-%m-%d 2015-09-28 +%s

path="/tmp/shell"
need_path="/tmp/logical"

if [ -e $path ]; then
    # 判断文件是否存在
    if [ -f $path ]; then
        rm $path
        if [ -d $need_path ]; then
            rm -rf $need_path
        else
            mkdir $need_path
        fi
    fi
else
    touch $path
fi


etc_pwd="/etc/passwd"
cat $etc_pwd | grep -v "^#" |awk '{FS=":"} {printf "The " NR "account is " $1 "\n"}'
