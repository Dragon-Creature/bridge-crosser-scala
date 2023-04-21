install: lint
	npm run build
	go build -o app ./cmd

lint:
	npm run prettier:fix

run:
	go run ./cmd

run-frontend:
	npm run start
