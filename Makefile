full-deploy: deploy dockerize
build: main
	go build main.go
deploy: 
	helm upgrade --install blog \
		--create-namespace -n blog \
		--set pullPolicy=Never \
		./helm
dockerize:
	docker build . -t app:latest
