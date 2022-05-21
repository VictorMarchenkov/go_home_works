doc:
	echo: "Starting generate documentation"
	godoc -http=localhost:6060

win:
	echo: "Build for Windows OS"
	go build -o bin/windows/wadding.exe cmd/app/main.go

unix:
	echo: "Build for FreeBSD OS"
	gox -osarch="freebsd/amd64" -output ./bin/unix/fadding ./cmd/app

linux:
	echo: "Build for Linux OS"
	gox -osarch="linux/amd64" -output ./bin/linux/ladding ./cmd/app


test:
	echo: "Start test"
	go test ./internal/controller -v