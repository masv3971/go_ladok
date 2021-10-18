.PHONY: update clean build build-all run package deploy test authors dist check-tag

NAME 					:= goladok3
VERSION                 := $(shell cat VERSION)
TAGS					:= $(shell git tag)

default: release-patch

release-patch: check-tag check-files tidy test add commit release-tag push-tag go-list push-main
		$(info relese ${NAME}@${VERSION})

tidy:
		$(info tidy up..)
		go mod tidy

test:
		$(info test ${NAME})
		go test -v --cover .

git-status:
	$(info files to be added:)
	@git status
	read -p "Press enter in order to precede"

add: git-status
	git add .

commit:
		git commit -S -m"${NAME} release ${VERSION}"

release-tag:
		git tag ${VERSION}

push-tag:
		git push origin ${VERSION}

push-main:
		git push origin main

check-tag:
	git fetch --tags
ifeq ($(filter $(TAGS), $(VERSION)) ,$(VERSION))
	$(error $(VERSION) is already used, make other one please)
endif

go-list:	
		GOPROXY=proxy.golang.org go list -m github.com/masv3971/${NAME}@${VERSION}

check-files: check-version-file check-license-file check-readme-file

check-version-file:
ifeq (,$(wildcard ./VERSION))
	$(error version file does not exists, make it!)
endif

check-license-file:
ifeq (,$(wildcard ./LICENSE.md))
	$(error license file does not exists, make it!)
endif

check-readme-file:
ifeq (,$(wildcard ./README.md))
	$(error README file does not exists, make it!)
endif
