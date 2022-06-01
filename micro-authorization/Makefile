postgresql_image := postgres:14.2-alpine3.15
container_name := micro-authorization-db
postgresql_user := postgres
postgresql_password := secret
postgresql_database := authorization

migration_dir := db/migration

installpostgre:
	docker run -d --name ${container_name} \
	-p 5432:5432 \
	-e POSTGRES_USER=${postgresql_user} \
	-e POSTGRES_PASSWORD=${postgresql_password} \
	-e POSTGRES_DB=${postgresql_database} \
	${postgresql_image}

stoppostgre:
	docker container stop ${container_name}

startpostgre:
	docker container start ${container_name}

uninstallpostgre:
	docker container stop ${container_name}
	docker container rm ${container_name}

createmigrate:
	migrate create -ext sql -dir ${migration_dir} -seq init_schema

migrateup:
	migrate -path ${migration_dir} \
	-database "postgresql://${postgresql_user}:${postgresql_password}@localhost:5432/${postgresql_database}?sslmode=disable" \
	-verbose up

.PHONY: installpostgre stoppostgre startpostgre uninstallpostgre createmigrate migrateup