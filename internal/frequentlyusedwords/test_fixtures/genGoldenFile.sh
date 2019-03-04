#!/bin/bash
cat mobydick.txt | tr -cs 'a-zA-Z' '[\n*]' | grep -v "^$" | tr '[:upper:]' '[:lower:]'| sort | uniq -c | sort -nr |
head -20 >> golden.out
