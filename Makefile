.PHONY: build

build:
	GOARM=7 GOARCH=arm GOOS=linux go build -o gopi_linux_arm7

deploy:
	scp -r ./public gopi_linux_arm7 gopi@192.168.1.33:/home/gopi/

statics:
	scp -r ./public gopi@192.168.1.33:/home/gopi/