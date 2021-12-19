.PHONY: dev
dev:
	${GOPATH}/bin/CompileDaemon \
		-exclude-dir=".git" \
		-exclude-dir=data -exclude=".#*" \
		-recursive=true \
	 	-build="go build -o ./main ./cmd/main.go" \
	 -command="./main"