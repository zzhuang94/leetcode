#!/usr/bin/env bash

# 194. Transpose File
# 
# Given a text file file.txt, transpose its content.
# 
# You may assume that each row has the same number of
# columns and each field is separated by the ' ' character.
# 
# For example, if file.txt has the following content:
# 
# name age
# alice 21
# ryan 30
# Output the following:
# 
# name alice ryan
# age 21 30

echo "name age no
alice 21 23
kobe 8 23s
ryan 30 24" > file.txt

# my solution Memory Limit Exceeded!!!!
#
# row=`sed -n "1p" file.txt| wc -w`
# 
# for i in $(seq 1 $row)
# do
#     cut -d ' ' -f $i file.txt | xargs
# done

awk '  
{  
    for(i=1;i<=NF;i++){  
        if(NR==1){  
            s[i]=$i;  
        }else{  
            s[i]=s[i]" "$i;  
        }  
    }  
}    
END{  
    for(i=1;s[i]!="";i++)  
        print s[i];  
}  
' file.txt 
