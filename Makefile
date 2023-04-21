install: lint
	npm run build
	go build -o app ./cmd

lint:
	npm run prettier:fix

run:
	go run ./cmd

run-frontend:
	npm run start

test:
	go test ./...
	npm run test

deploy: install
	tar -czvf deploy.tar.gz build app
	scp deploy.tar.gz ubuntu@3.92.200.64:deploy.tar.gz
	ssh ubuntu@3.92.200.64 rm -r build
	ssh ubuntu@3.92.200.64 rm app
	ssh ubuntu@3.92.200.64 tar -xvf deploy.tar.gz