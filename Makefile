IMAGE_NAME=glog-authentication
VERSION_FILE=VERSION
IMAGE_VERSION=$(shell cat VERSION)

test-build: release-build dockerize

release-build:
	./ci/semver.sh $(VERSION_FILE) release-build

release-patch:
	./ci/semver.sh $(VERSION_FILE) release-patch

release-minor:
	./ci/semver.sh $(VERSION_FILE) release-minor

release-major:
	./ci/semver.sh $(VERSION_FILE) release-major

build: main
	go build main.go

test:
	go test ./...

dockerize:
	docker build . -t ${IMAGE_NAME}:${IMAGE_VERSION}
