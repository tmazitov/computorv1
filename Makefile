NAME	= computorv1

all:
	go build -o $(NAME) .

clean:
	go clean

fclean: clean
	rm -f $(NAME)

re: fclean all

.PHONY: all clean fclean re
