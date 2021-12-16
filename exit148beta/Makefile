# Run beta

run:
	go build -o adminportalbeta ../main.go ../server.go
	sudo install -d -m 0775 /usr/local/adminportalbeta/bin
	sudo install -m 0755 -o 0 -g 0 adminportalbeta         /usr/local/adminportalbeta/bin/adminportalbeta
	sudo install -m 0755 -o 0 -g 0 rc.beta.FreeBSD /usr/local/etc/rc.d/adminportalbeta
	( cd .. && tar -c -f - public $$(find routes -type f -name '*.html' ) ) | tar xpvf -
	-sudo /usr/local/etc/rc.d/adminportalbeta stop
	sudo /usr/local/etc/rc.d/adminportalbeta start
	-sudo /usr/local/etc/rc.d/nginx reload
