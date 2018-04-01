#!/usr/bin/make
 
.DEFAULT_GOAL := all

DIST_DIR := dist
PLATFORMS := linux/amd64 linux/386 darwin/amd64 darwin/386 windows/386 windows/amd64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

releases: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build $(LD_FLAGS) -o 'dist/kibup_$(os)-$(arch)'

release:
	go build $(LD_FLAGS) -o 'dist/kibup'

clean:
	@rm -rf $(DIST_DIR)

all:
	@make -s clean release