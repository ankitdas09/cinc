executor:
	docker build -t executor . && docker run -v /var/run/docker.sock:/var/run/docker.sock -v /tmp:/tmp -it --rm -p 8000:8000 executor:latest