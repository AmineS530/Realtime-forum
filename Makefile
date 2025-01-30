NAME=Real-time-forum

all: $(NAME) run

$(NAME):
	go build -o $(NAME).exec main.go

run: $(NAME)
	./$(NAME).exec

clean:
	rm -f $(NAME).exec

re: clean all

.PHONY: all clean