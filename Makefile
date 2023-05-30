build:
	go build -o bin/readFromFile .

test: build
	test/test.sh

build-image:
	docker build -t readfromfile:1 .

test-image: build-image
	test/test-image.sh