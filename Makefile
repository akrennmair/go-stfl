include $(GOROOT)/src/Make.inc

TARG=stfl
CGOFILES=stfl.go
CLEANFILES+=example

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$(O)

