# readFromFile

ReadFromFile is a sample golang program, to let the user read both from file or from the standard input.

The program reads an input of json array of objects. Each object contains two string fields, `name` and 
`email`. for example:
```json
[
  {
    "name": "tester1",
    "email": "tester1.example.com"
  },
  {
    "name": "tester2",
    "email": "tester2.example.com"
  }
]
```
The output is a textual format of the same data, e.g.
```text
Name: tester1
email: tester1.example.com
Name: tester2
email: tester2.example.com
```
## Build
To build the program, run 
``` shell
make build
```

To build a docker image, run
``` shell
make build-image
```

## Usage
There are two options to run this program
### Get input from the standard input
To get input from the standard input, run the program without any parameter. Then type a 
valid json content. For example:
```shell
$ bin/readFromFile
[{"name":"example","email":"example@example.com"}]
```
Or use pipe to send data to the standard input; e.g.
```shell
cat <<EOF | bin/readFromFile
[
  {
    "name": "tester1",
    "email": "tester1.example.com"
  },
  {
    "name": "tester2",
    "email": "tester2.example.com"
  }
]
EOF
```
Another example:
```shell
cat testdata/test2.json | bin/readFromFile
```
### Send file name(s) as argument(s)
For example
```shell
bin/readFromFile testdata/test1.json
```
or 
```shell
bin/readFromFile testdata/test1.json testdata/test2.json
```

or even
```shell
bin/readFromFile testdata/*.json
```


## Testing
To test the program, run
``` shell
make test
```

To test the image, run
``` shell
make test-image
```
