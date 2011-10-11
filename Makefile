include $(GOROOT)/src/Make.inc

TARG=github.com/akrennmair/go-stfl
CGOFILES=stfl.go
CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$(O)

