#! /bin/bash
app=stripe
which $app

if [ $? -ne 0 ]; then
  echo "$app command not found. Please check your environment PATH."
  exit 1
fi

cwd="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
dir="$cwd/data"

files=($(ls -d $dir/* | grep -i "\d.json$"))

length=${#files[@]}

for ((i=0; i<$length; i++)); do
   file=${files[$i]}
   num=$((i + 1))
   echo "INPUT $num:"
   cat $file | jq .
   echo ""
   result=$($app part1 -f $file)
   echo "OUTPUT $num:"
   echo $result | jq .
done
