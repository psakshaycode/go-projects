.PHONY: postgres adminer migrate

postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres

adminer:
	docker run --rm -ti --network host adminer

migration:
	migrate -source file://migrations \
					-database postgres://postgres:secret@localhost/postgres?sslmode=disable up