install: lint
	npm run build
	go build -o app ./cmd

lint:
	npm run prettier:fix

dependencies:
	go mod tidy
	go install github.com/matryer/moq@latest
	npm install

run:
	npm run build
	go run ./cmd

run-frontend:
	npm run start

test:
	go test ./...
	npm run test

deploy: install test
	tar -czvf deploy.tar.gz build app bridge-crosser-scala.service
	scp deploy.tar.gz ubuntu@3.92.200.64:deploy.tar.gz
	ssh ubuntu@3.92.200.64 sudo rm -rf /opt/bridge-crosser-scala
	ssh ubuntu@3.92.200.64 sudo mkdir /opt/bridge-crosser-scala
	ssh ubuntu@3.92.200.64 sudo tar -xvf deploy.tar.gz -C /opt/bridge-crosser-scala
	ssh ubuntu@3.92.200.64 sudo mv /opt/bridge-crosser-scala/bridge-crosser-scala.service /etc/systemd/system
	ssh ubuntu@3.92.200.64 sudo systemctl daemon-reload
	ssh ubuntu@3.92.200.64 sudo systemctl start bridge-crosser-scala.service
	ssh ubuntu@3.92.200.64 sudo systemctl enable bridge-crosser-scala.service