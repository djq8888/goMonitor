#需要在crontab中配置每分钟运行1次本脚本
#* * * * * sh /path/processMonitor.sh process_name &
#根据进程名获取进程号
pid=$(ps -e | grep $1 | awk '{print $1}')
#创建两个文件.monitor文件用于记录所有信息，.latest文件由于记录最近1小时的信息
touch $1.monitor $1.latest
#每隔5s获取一次，只保留CPU和内存占用信息，每分钟写入.monior文件
top -n 12 -d 5 -p $pid -b | grep $1 | awk '{print $9,$10}' >> $1.monitor
#将最近1小时的数据写入.latest文件
tail -n 612 $1.monitor > $1.latest