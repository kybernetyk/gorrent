include $(GOROOT)/src/Make.inc

DEPS=\
		 bencode 

TARG=gorrent
GOFILES=\
	gorrent.go

include $(GOROOT)/src/Make.cmd
