#统计重复出现的字符串
echo $1
cat $1 | awk '{print $1}' | awk -F: '{arr[$1]++} END {for (i in arr) print i, arr[i]}' | sort -rnk2

#bash repeat.sh  repeat.txt (统计样本)  

