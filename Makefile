appsrv:
ifeq ($(action),up)
	@ cd services/appsrv && docker-compose up -d
else
	@ cd services/appsrv && docker-compose down
endif

contactsrv:
ifeq ($(action),up)
	@ cd services/contactsrv && docker-compose up -d
else
	@ cd services/contactsrv && docker-compose down
endif

todosrv:
ifeq ($(action),up)
	@ cd services/todosrv && docker-compose up -d
else
	@ cd services/todosrv && docker-compose down
endif