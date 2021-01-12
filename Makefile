MAKE=make

DEP=dep ensure -v
all:
	+$(DEP)
	+$(MAKE) -C client/
	+$(MAKE) -C server/

clean:
	rm -f client/client_svc
	rm -f server/server_svc