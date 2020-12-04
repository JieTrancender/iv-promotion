BUILDDIR = build
BUILDFLAGS =


.PHONY: all clean test

APPS = atomic-test nsq-producer
all: $(APPS)

clean:
	rm -rf $(BUILDDIR)

$(BUILDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${BUILDFLAGS} -o $@ ./apps/$*

.PHONY: $(APPS)
$(APPS): %: $(BUILDDIR)/%

test:
	go test -v -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
