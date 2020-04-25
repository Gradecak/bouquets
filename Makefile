.PHONY: build sample

DIR := ${CURDIR}

build:
	docker build -t bouquets -f build/Dockerfile .

sample: build
	docker run -v ${DIR}/sample.txt:/sample.txt bouquets:latest /sample.txt
