#!/usr/bin/env bash
set -ex

mkdir -p _out

bin/readFromFile testdata/test1.json testdata/test2.json > _out/res1.txt
diff testdata/expected.txt _out/res1.txt

cat testdata/test1.json | bin/readFromFile > _out/res2.txt
bin/readFromFile < testdata/test2.json >> _out/res2.txt
diff testdata/expected.txt _out/res2.txt
