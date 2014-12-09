#!/bin/sh
BASEPATH=$(cd `dirname $0`;pwd)
ADDRESS=${BASEPATH}/backup
DATABASE=FishMoney

#add data
DATE=`date "+%Y-%m-%d_%T"`
/usr/bin/mysqldump -uroot -p1 ${DATABASE} > ${ADDRESS}/$DATE

#del data
LASTSAVEDATE=`date -d last-month "+%Y-%m-%d_%T"`
for f in `ls ${ADDRESS} | xargs echo`
	do
	if [ "${f}" \< "$LASTSAVEDATE" ];then
		rm ${ADDRESS}/${f}
	fi
done
