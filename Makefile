default:
	go build -o news-scraper-macos
	clear
	./news-scraper-macos
github:
	@echo "Committing changes to Github..."
	git add -A
	git commit -m "$m"
	git push
arm-build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o news-scraper-arm
