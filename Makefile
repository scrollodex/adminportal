
all:
	cd exit148beta && make

.PHONY: port
port: exit148prod/.env-NEW

.PHONY: exit148prod/.env-NEW
exit148prod/.env-NEW:
	sed <exit148beta/.env >exit148prod/.env-NEW  \
	-e 's/beta-scrollodex-db/scrollodex-db/g' \
	-e 's/prod-scrollodex-db/scrollodex-db/g' \
	-e 's/beta/prod/g'  \
	-e 's/3000/3001/g' 

.PHONY: reset-test-data
reset-test-data:
	bin/reset-test-data
