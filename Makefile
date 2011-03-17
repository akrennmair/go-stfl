include $(GOROOT)/src/Make.inc

TARG=stfl
CGOFILES=stfl.go
CLEANFILES+=example

CGO_LDFLAGS:=`pkg-config --static --libs stfl`
CGO_CFLAGS:=`pkg-config --cflags stfl`

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$(O)

