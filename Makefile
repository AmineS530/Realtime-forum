NAME=Real-time-forum

all: $(NAME)

$(NAME):
# go build -o $(NAME).exec main.go
	@go run main.go

#run: $(NAME)
#	@./$(NAME).exec

clean:
	@echo "Deleteing the DataBase and executable"
	@rm -f $(NAME).exec
	@rm -f DataBase/RTF.db

re: clean all

.PHONY: all clean