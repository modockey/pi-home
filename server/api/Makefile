generate:
	java -jar openapi-generator-cli.jar generate \
    -i ./swagger.yaml \
    -g rust-server \
    -o .

.PHONY: release-build
release-build:
	cargo build --release

.PHONY: release-systemd
release-systemd:
	sudo cp ./systemd/home-management-api.service /etc/systemd/system/

release: release-build release-systemd
	systemctl stop home-management-api.service
	sudo systemctl daemon-reload
	sudo mkdir -p /usr/local/bin/home-management-api
	sudo cp ./target/release/openapi_client /usr/local/bin/home-management-api/
	systemctl start home-management-api.service
	systemctl enable home-management-api.service
