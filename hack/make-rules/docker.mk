export DOCKER_USERNAME ?=
export DOCKER_PASSWORD ?=
export DOCKER_HOSTNAME ?= ghcr.io
export DOCKER_NAMESPACE ?= revit13
export DOCKER_TAGNAME ?= master
DOCKER_CONTEXT ?= .

DOCKER_FILE ?= Dockerfile
DOCKER_NAME ?= vault-plugin-secrets-kubernetes-reader

IMG ?= ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/${DOCKER_NAME}:${DOCKER_TAGNAME}


.PHONY: docker-build
docker-build:
	docker build $(DOCKER_CONTEXT) -t ${IMG} -f $(DOCKER_FILE)

.PHONY: docker-push
docker-push:
ifneq (${DOCKER_PASSWORD},)
	@docker login \
		--username ${DOCKER_USERNAME} \
		--password ${DOCKER_PASSWORD} ${DOCKER_HOSTNAME}
endif
	docker push ${IMG}

.PHONY: docker-rmi
docker-rmi:
	docker rmi ${IMG} || true

