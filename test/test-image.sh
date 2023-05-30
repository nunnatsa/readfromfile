#!/usr/bin/env bash
set -ex

mkdir -p _out

docker run --rm -v $(pwd)/testdata:/testdata:Z readfromfile:1 testdata/test1.json testdata/test2.json > _out/res1.txt
diff testdata/expected.txt _out/res1.txt

cat testdata/test1.json | docker run -i --rm readfromfile:1 > _out/res2.txt
docker run -i --rm readfromfile:1 < testdata/test2.json >> _out/res2.txt
diff testdata/expected.txt _out/res2.txt
