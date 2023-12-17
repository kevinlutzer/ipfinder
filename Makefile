REPO=klutzer/ipfinder
VERSION=latest


start: 
	API_KEY=asd go run main.go

build:
	docker buildx build --platform linux/amd64 --tag $(REPO):$(VERSION) .

push: 
	docker push $(REPO):$(VERSION)