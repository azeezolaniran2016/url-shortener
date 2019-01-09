APP_NAME := "url-shortener"

.PHONY: dev/server
dev/server:
	@echo "Starting $(APP_NAME) server container"
	@docker-compose -f docker-compose.loc.yml up --build

.PHONY: dev/clean-server
dev/clean-server:
	@echo "Removing server container, images and volumees"
	@docker-compose -f docker-compose.loc.yml down --rmi all

.PHONY: gin/server
gin/server:
	@gin -i -a 80 run main.go



