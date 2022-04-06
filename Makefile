build:
	go build -o main .

docker:
	@echo "Docker image build:"
	docker image build -f Dockerfile -t forum-image .
	@echo

	@echo "Docker images:"
	docker images
	@echo

	@echo "Container run:"
	docker container run -t -p 8080:8080 --detach --name forum-container forum-image
	@echo

	@echo "Server run:"
	docker exec -it forum-container ./main
	@echo

stop:
	@echo "Container stop:"
	docker stop forum-container
	@echo

clean:
	@echo "Container remove:"
	docker rm forum-container
	@echo

	@echo "Deleting images:"
	docker rmi -f forum-image
	@echo

	@echo "List of images and containers now:"
	docker ps -a
	@echo
	docker images
	@echo

	rm main

.DEFAULT_GOAL := build