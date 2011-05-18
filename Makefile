include $(GOROOT)/src/Make.inc

DEPS=\
		 bencode 

TARG=gorrent
GOFILES=\
	rfc1738.go\
	metainfo.go\
	gorrent.go

include $(GOROOT)/src/Make.cmd
