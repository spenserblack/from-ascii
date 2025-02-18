from-ascii: *.go
	go build -o from-ascii ./...

.PHONY: install
install: from-ascii
	cp from-ascii /usr/local/bin/from-ascii
