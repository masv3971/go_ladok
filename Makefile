.PHONY: update clean build build-all run package deploy test authors dist check-tag

NAME 					:= goladok3
VERSION                 := $(shell cat VERSION)
TAGS					:= $(shell git tag)

default: release-patch

release-patch: check-tag tidy test add commit-msg release-tag push-tag go-list
		@echo relese ${NAME}@${VERSION} 

tidy:
		@echo tidy up..
		go mod tidy

test:
		@echo test ${NAME}
		go test -v --cover .
add:
	git add .

commit-msg:
		git commit -S -m"${NAME} release ${VERSION}"

release-tag:
		git tag ${VERSION}

push-tag:
		git push origin ${VERSION}

check-tag:
ifeq ($(filter $(TAGS), $(VERSION)) ,$(VERSION))
	@echo tag $(VERSION) is already used, make other one please
	@exit 
endif

go-list:	
		GOPROXY=proxy.golang.org go list -m github.com/masv3971/${NAME}@${VERSION}
