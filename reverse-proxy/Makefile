.PHONY: release
release: ./reverse_proxy.conf
	sudo cp ./reverse_proxy.conf /etc/nginx/conf.d/
	sudo cp ./ssl.conf /etc/nginx/conf.d/
	sudo systemctl restart nginx
