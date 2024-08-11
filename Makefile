build:
	go build -o wechat-router .

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o wechat-router .

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o wechat-router .

build-window:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o wechat-router .