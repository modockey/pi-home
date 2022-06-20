.PHONY: release-build
release-build:
	cargo build --release

.PHONY: release-systemd
release-systemd:
	cp ./systemd/ip_register.service /etc/systemd/system/
	cp ./systemd/ip_register.timer /etc/systemd/system/

release: release-build release-systemd
	mkdir -p /usr/local/bin/ip_register
	cp ./target/release/ip_register /usr/local/bin/ip_register/
	systemctl enable ip_register.timer
