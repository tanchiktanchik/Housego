test:
	@echo 'Мы сделали Makefile'

up:
	sudo -S docker-compose up --build awesome-project

stop:
	sudo -S docker-compose stop
