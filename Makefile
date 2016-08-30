DESTINATION=_dst/fluka
default: build

build:
	go build -o $(DESTINATION)

