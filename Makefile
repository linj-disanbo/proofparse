.PHONY: build package clean

build:
	@go build -o ./parser ./cmd/cmd.go

package:
	@mkdir -p parser-bin
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o ./parser-bin/parser_darwin_amd64 ./cmd/cmd.go
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./parser-bin/parser_linux_amd64 ./cmd/cmd.go
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ./parser-bin/parser_windows_amd64.exe ./cmd/cmd.go
	@tar -zcvf ./parser-bin.tar.gz ./parser-bin

clean:
	@rm -rf ./parser
	@rm -rf ./parser-bin*
