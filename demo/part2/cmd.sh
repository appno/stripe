#! /bin/bash

DIR="$PWD/data"
FILES=$(ls -d $DIR/* | grep -i "\d.json$")
APP=stripe
DELAY=$1
DEADLINE=$2
UNITS="s"
STRIPE_DEADLINE=$DEADLINE$UNITS

export STRIPE_DEADLINE=$STRIPE_DEADLINE

which $APP

if [ $? -ne 0 ]; then
  echo "$APP command not found"
  exit 1
fi

for FILE in ${FILES[@]}
do
   echo -e "$FILE\n"
   $APP part2 -f $FILE
   sleep $DELAY
done
