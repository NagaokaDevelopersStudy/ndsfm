VERSION=0.0.2
PROJECT=ndsfm
APP=pingpong

deploy: deployment.yaml
	kubectl apply -f deployment.yaml

push: docker_build
	docker tag $(PROJECT)/$(APP):$(VERSION) gcr.io/$(PROJECT)/$(APP):$(VERSION)
	gcloud --project=$(PROJECT) docker -- push gcr.io/$(PROJECT)/$(APP):$(VERSION)

docker_build: Dockerfile entrypoint.sh pingpong.alpine pingpong.json
	docker build -t $(PROJECT)/$(APP):$(VERSION) .

pingpong.alpine: main.go Gopkg.lock
	GOOS=linux GOARCH=amd64 go build -o pingpong.alpine
	upx pingpong.alpine

Gopkg.lock: Gopkg.toml
	dep ensure
