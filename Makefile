isDocker := $(shell docker info > /dev/null 2>&1 && echo 1)

isContainerRunning := $(shell docker ps | grep gocord > /dev/null 2>&1 && echo 1)
user := $(shell id -u)
group := $(shell id -g)

ifeq ($(isDocker), 1)
    ifeq ($(isContainerRunning), 1)
        dc := USER_ID=$(user) GROUP_ID=$(group) docker-compose
        ge := docker exec -u $(user):$(group) go
        gr := docker exec -u $(user):$(group) go run
    else
        dc := USER_ID=$(user) GROUP_ID=$(group) docker-compose
        ge :=
    endif
else
    ge :=
endif

build-docker:
	$(dc) pull --ignore-pull-failures
	$(dc) build --no-cache

up:
	$(dc) up -d

stop:
	$(dc) stop

main:
	$(gr) main.go