Count=`wc -l zzz.log | awk '{print $1}'`
echo "Count is $Count"
num=1
while(($num<=$Count))
do
	sed -n ${num}p zzz.log >> mysql-slow.log
	echo $num
	let num=$num+1
	sleep 1
done
