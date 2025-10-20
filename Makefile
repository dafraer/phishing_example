run:
	docker build --no-cache -t phishing-example:1.0 . && docker compose up -d
stop:
	docker compose down