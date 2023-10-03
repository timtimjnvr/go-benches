#!/bin/zsh
export LC_ALL=C
for i in {0..10000..1000}
    do
        cat /dev/urandom |tr -dc '[:alpha:]'|fold -w ${1:-$i} | head -n1 >> $i.txt
done
