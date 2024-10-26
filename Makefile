.PHONY: protoc-go build-and-push-image build-local build-remote install run

protoc-go:
	./script/protoc.sh

build-and-push-image:
	./script/build-and-push-image.sh

dev1-deploy:
	./script/deploy-latest.sh --env

build-local:
	./script/build_local.sh

install:
	./script/install.sh

run:
	./script/run.sh