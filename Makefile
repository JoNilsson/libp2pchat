help:
	@echo "IllChat Makefile (Requires Go v1.16+)"
	@echo "'help' - Displays the command usage"
	@echo "'build' - Builds the application into a binary/executable" 
	@echo "'install' - Installs the application"
	@echo "'build-windows' - Builds the application for Windows platforms"
	@echo "'build-darwin' - Builds the application for MacOSX platforms"
	@echo "'build-linux' - Builds the application for Linux platforms"
	@echo "'build-all' - Builds the application for all platforms"

build:
	@echo Compiling IllChat
	@go build .
	@echo Compile Complete. Run './illchat_v1(.exe)'

install:
	@echo Installing IllChat
	go install .
	@echo install Complete. Run 'illchat_v1'.

build-windows:
	@echo Cross Compiling IllChat for Windows x86
	@GOOS=windows GOARCH=386 go build -o ./bin/illchat_v1-windows-x32.exe
	@echo Cross Compiling IllChat for Windows x64
	@GOOS=windows GOARCH=amd64 go build -o ./bin/illchat_v1-windows-x64.exe

build-darwin:
	@echo Cross Compiling IllChat for MacOSX x64
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/illchat_v1-darwin-x64

build-linux:
	@echo Cross Compiling IllChat for Linux x32
	@GOOS=linux GOARCH=386 go build -o ./bin/illchat_v1-linux-x32
	@echo Cross Compiling IllChat for Linux x64
	@GOOS=linux GOARCH=amd64 go build -o ./bin/illchat_v1-linux-x64
	@echo Cross Compiling IllChat for Linux Arm32
	@GOOS=linux GOARCH=arm go build -o ./bin/illchat_v1-linux-arm32
	@echo Cross Compiling IllChat for Linux Arm64
	@GOOS=linux GOARCH=arm64 go build -o ./bin/illchat_v1-linux-arm64

build-all: build-windows build-darwin build-linux
	@echo Cross Compiled IllChat for all platforms

