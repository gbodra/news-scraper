default:
	go build -o news-scraper-macos
	clear
	./news-scraper-macos
arm-build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o news-scraper-arm
