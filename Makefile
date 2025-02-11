NAME=Real-time-forum

all: $(NAME)

$(NAME):
# go build -o $(NAME).exec main.go
	@go run main.go

#run: $(NAME)
#	@./$(NAME).exec

clean:
	@rm -f $(NAME).exec

re: clean all

.PHONY: all clean