APP_NAME=main.exe

build:
	go build -o $(APP_NAME)
	./$(APP_NAME)

clean:
	rm -f $(APP_NAME)

re: clean build