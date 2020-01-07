appsrv:
ifeq ($(action),up)
	@ cd services/appsrv && docker-compose up -d
else
	@ cd services/appsrv && docker-compose down
endif


todosrv:
ifeq ($(action),up)
	@ cd services/todosrv && docker-compose up -d
else
	@ cd services/todosrv && docker-compose down
endif