BUILDDIR = build
BUILDFLAGS =


.PHONY: all clean

APPS = atomic-test
all: $(APPS)

clean:
	rm -rf $(BUILDDIR)

$(BUILDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${BUILDFLAGS} -o $@ ./$*

.PHONY: $(APPS)
$(APPS): %: $(BUILDDIR)/%
