install:
	@echo "Please specify a target to install to: nginx, apache [root/sudo required]"

nginx:
	./scripts/install_nginx.sh
apache:
	./scripts/install_apache.sh


