# Makefile

build:
	go build -o config-updater main.go

clean:
	rm -f ./config-updater
