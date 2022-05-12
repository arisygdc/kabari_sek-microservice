postgresql_image := postgres:14.2-alpine3.15
container_name := micro-user-db
postgresql_user := postgres
postgresql_password := secret

installpostgre:
	docker run -d --name ${container_name} \
	-p 5432:5432 \
	-e POSTGRES_USER=${postgresql_user} \
	-e POSTGRES_PASSWORD=${postgresql_password} \
	${postgresql_image}

stoppostgre:
	docker container stop ${container_name}

startpostgre:
	docker container start ${container_name}

uninstallpostgre:
	docker container stop ${container_name}
	docker container rm ${container_name}

.PHONY: installpostgre stoppostgre startpostgre uninstallpostgre