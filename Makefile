name = "docv"
project_path = "./cmd/$(name)/main.go"

run:
	go build -o $(name) $(project_path)
	./$(name)
