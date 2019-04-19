test:
	GO111MODULE=on go test ./...

reset-mod:
	git checkout go.mod go.sum

