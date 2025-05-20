NAME=Real-Time-Forum

ENV_FILE=.env

PORT=$(firstword $(filter-out all build run clean,$(MAKECMDGOALS)))

PORT_MSG := $(if $(PORT),port: $(PORT),default port)

all: $(ENV_FILE) $(NAME)

$(PORT): $(NAME)

$(NAME): $(ENV_FILE)
	@echo "\033[1;32mRunning $(NAME) on $(PORT_MSG)\033[0m"
	@PORT=$(PORT) go run .

$(ENV_FILE):
	@echo "Checking for .env file..."
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Creating .env file with default SECRET_KEY..."; \
		echo 'SECRET_KEY=#Forum@zone01!' > $(ENV_FILE); \
	else \
		echo ".env file already exists."; \
	fi

clean:
	@echo "Deleteing the DataBase and executable"
	@rm -f $(NAME).exec
	@rm -f DataBase/RTF.db

re: clean all

.PHONY: all clean