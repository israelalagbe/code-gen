CWD := "$(abspath $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))"
run:
	
	go run cmd/code-gen/main.go $(CWD) /Users/israelalagbe/Projects/experiment
	
build:
	cp -r ./templates ./bin
	go build -o bin/code-gen cmd/code-gen/main.go