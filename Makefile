db_login:
	psql ${url}

migrate:
	migrate create -ext sql -dir package/migrations -seq ${name}

migrate_up:
	migrate -database ${url} -path package/migrations up

migrate_down:
	migrate -database ${url} -path package/migrations down

migrate_drop:
	migrate -database ${url} -path package/migrations drop -f