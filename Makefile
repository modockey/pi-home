release-all:
	make -C ./nginx release
	make -C ./server/ip_register release
	make -C ./server/api release
