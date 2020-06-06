default:
	go build && clear
arm-build:
	env GOOS=linux GOARCH=arm GOARM=5 go build
