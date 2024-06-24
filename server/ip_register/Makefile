.PHONY: release-build
release-build:
	cargo build --release

.PHONY: release-systemd
release-systemd:
	sudo cp ./systemd/ip_register.service /etc/systemd/system/
	sudo cp ./systemd/ip_register.timer /etc/systemd/system/

release: release-build release-systemd
	sudo mkdir -p /usr/local/bin/ip_register
	sudo cp ./target/release/ip_register /usr/local/bin/ip_register/
	sudo systemctl start ip_register.service
	sudo systemctl enable ip_register.service
	sudo systemctl start ip_register.timer
	sudo systemctl enable ip_register.timer
