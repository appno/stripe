#! /bin/bash
if [ $# != 2 ]; then
    echo "Usage: $0 [DELAY] [DEADLINE]"
    exit 1
fi

app=stripe
which $app

if [ $? -ne 0 ]; then
  echo "$app command not found. Please check your environment PATH."
  exit 1
fi

cwd="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
dir="$cwd/data"

files=($(ls -d $dir/* | grep -i "\d.json$"))
delay=$1
deadline=$2
units="s"
stipe_deadline=$deadline$units
port=8083
stripe_home=`mktemp -d 2>/dev/null || mktemp -d -t 'stripe_home'`

export STRIPE_DEADLINE=$stipe_deadline
export STRIPE_HOME=$stripe_home

echo "STRIPE_HOME=$STRIPE_HOME"
echo -n "Starting server on localhost:$port"
$app server $port &>/dev/null &
pid=$!
# echo $pid > $stripe_home/server.pid
sleep 2
echo -e "\rServer started on localhost:$port  "

length=${#files[@]}
time=0
suffix="s"

for ((i=0; i<$length; i++)); do
   file=${files[$i]}
   num=$((i + 1))
   echo "Time $time$suffix:"
   echo "REQUEST $num:"
   cat $file | jq .
   echo ""
   response=$(curl -s -H "Content-Type: application/json" -d "@$file" localhost:$port)
   echo "RESPONSE $num:"
   echo $response | jq .
   if (( $i != $length - 1)); then
     sleep $delay
   fi
done

echo -n "Shutting down server..."
kill $pid
echo -e "\rServer shut down successfully."
