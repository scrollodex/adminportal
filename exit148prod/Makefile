# Run prod

run:
	go build -o adminportalprod ../main.go ../server.go
	sudo install -d -m 0775 /usr/local/adminportalprod/bin
	sudo install -m 0755 -o 0 -g 0 adminportalprod         /usr/local/adminportalprod/bin/adminportalprod
	sudo install -m 0755 -o 0 -g 0 rc.prod.FreeBSD /usr/local/etc/rc.d/adminportalprod
	( cd .. && tar -c -f - public $$(find routes -type f -name '*.html' ) ) | tar xpvf -
	-sudo /usr/local/etc/rc.d/adminportalprod stop
	sudo /usr/local/etc/rc.d/adminportalprod start
	-sudo /usr/local/etc/rc.d/nginx reload
