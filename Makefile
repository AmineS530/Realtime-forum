NAME=Real-time-forum
ENV_FILE=.env

all: $(ENV_FILE) $(NAME)

$(NAME):
# go build -o $(NAME).exec main.go
	@go run main.go


$(ENV_FILE):
	@echo "Checking for .env file..."
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Creating .env file with default SECRET_KEY..."; \
		echo 'SECRET_KEY=#Forum@zone01!' > $(ENV_FILE); \
	else \
		echo ".env file already exists."; \
	fi

#run: $(NAME)
#	@./$(NAME).exec

clean:
	@echo "Deleteing the DataBase and executable"
	@rm -f $(NAME).exec
	@rm -f DataBase/RTF.db

re: clean all

.PHONY: all clean