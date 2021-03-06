BUILDDIR = build
BUILDFLAGS =

APPS = atomic-test nsq-producer nsq-consumer cobra-test zap-test waitGroup-test goroutine-test select-test \
	closeChannelToBroadcast multiGenerate chain-test multiTask multiTaskInFixed future-test context-test \
	producer-consumer publish-subscribe etcd-test pflag-test
all: $(APPS)

$(BUILDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${BUILDFLAGS} -o $@ ./apps/$*

$(APPS): %: $(BUILDDIR)/%

clean:
	rm -rf $(BUILDDIR)

test:
	GO111MODULE=on go test -v -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: clean all test lint
.PHONY: $(APPS)

lint:
	golangci-lint cache clean
	golangci-lint run --tests=false ./...
