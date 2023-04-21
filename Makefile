Install: Lint
	npm run build
	go build -o app ./cmd

Lint:
	npm run prettier:fix