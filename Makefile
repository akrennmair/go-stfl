include $(GOROOT)/src/Make.inc

TARG=github.com/akrennmair/go-stfl
CGOFILES=stfl.go

include $(GOROOT)/src/Make.pkg

example:
	$(MAKE) -C $@

.PHONY: example
