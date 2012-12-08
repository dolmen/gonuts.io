GOPATH:=$(shell pwd)/gopath
SDKROOT:=/usr/local/share/go-app-engine-64

GOFILES:=$(shell find . -name *.go)

all: fvb

prepare:
	rm -fr $(GOPATH)
	mkdir -p $(GOPATH)/src
	go get -u github.com/bmizerany/pat
	rm -fr $(GOPATH)/src/github.com/bmizerany/pat/.git
	rm -fr $(GOPATH)/src/github.com/bmizerany/pat/example
	nut get -v nut

# format, vet, build
fvb:
	gofmt -e -s -w .
	go tool vet .
	$(SDKROOT)/goroot/bin/go-app-builder -goroot=$(SDKROOT)/goroot -dynamic -unsafe $(GOFILES)

run: fvb
	dev_appserver.py --skip_sdk_update_check --use_sqlite .

run_clean: fvb
	dev_appserver.py --clear_datastore --use_sqlite .

deploy: fvb
	appcfg.py --oauth2 update .
