CONTAINER_NAME=glog
VERSION_FILE=VERSION
CONTAINER_VERSION=$(shell cat VERSION)

full-deploy: release-minor deploy dockerize

test-deploy: release-patch dockerize deploy

release-patch:
	./ci/semver.sh $(VERSION_FILE) release-patch

release-minor:
	./ci/semver.sh $(VERSION_FILE) release-minor

release-major:
	./ci/semver.sh $(VERSION_FILE) release-major

build: main
	go build main.go

deploy: 
	helm upgrade --install glog \
		--create-namespace -n glog \
		--set pullPolicy=Never \
		--set image.repository=${CONTAINER_NAME} \
		--set image.version=${CONTAINER_VERSION} \
		./ci/helm

dockerize:
	docker build . -t ${CONTAINER_NAME}:${CONTAINER_VERSION}
