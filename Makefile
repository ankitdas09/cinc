run:
	docker build -t compiler . && docker run -v /var/run/docker.sock:/var/run/docker.sock -it --rm compiler:latest