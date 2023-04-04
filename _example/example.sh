#!/bin/sh

O=1
n=$(ls *.html | wc -l)
m=4
p=$((m * 2))

rm -f *.csv
rm -f *.out
rm -f *.err
rm -f *.txt
rm -f *.tf
rm -f idf
rm -f *.tfidf

set -e

ih $n
tf $n
idf $n
tfidf $n
cs $O $n $m $p

set +e

cat l-*.csv > l.csv
cat m-*.csv > m.csv
cat h-*.csv > h.csv

rm -f *.txt
rm -f *.tf
rm -f idf
rm -f *.tfidf
rm -f l-*.csv
rm -f m-*.csv
rm -f h-*.csv
