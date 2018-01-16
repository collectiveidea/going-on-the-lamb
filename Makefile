build/deployment.zip: build/main
	cd build && zip deployment.zip main

build/main:
	GOOS=linux go build -o build/main

build:
	mkdir build

.PHONY: clean
clean:
	rm -rf build

.PHONY: create-function
create-function:
	aws lambda create-function \
		--region ${REGION} \
		--function-name HelloFunction \
		--zip-file fileb://./build/deployment.zip \
		--runtime go1.x \
		--tracing-config Mode=Active \
		--role ${ROLE_ARN} \
		--handler main

.PHONY: test-endpoint
test-endpoint:
	curl -XPOST -d "<your name here>" ${API_ENDPOINT}
