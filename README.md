# Frequently Used Words

Given a utf-8 text file as an argument, reads the file, and outputs the n most
frequently used words in the file in order, along with their frequency.​

## Usage
Make all

binary is in /bin

./frequentlyUsedWords:
  -file string
        text file you wish parse (default "internal/frequentlyusedwords/test_fixtures/mobydick.txt")

## Bench 
Target cat/tr/sort/uniq
real    0m0.118s
user    0m0.124s
sys     0m0.012s

Current state
real    0m1.370s
user    0m1.408s
sys     0m0.012s

## todo
see plan.txt for requirements

use binary-hash tree or other hash map
add bench and profile to improve time taken