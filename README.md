# run 

```
PORT=8080 go run main.go
curl -i http://localhost:8080/foo/bar
```

# create image

```
docker build http-echo .
docker run --rm -p 8080:80 --name http-echo http-echo
curl -i http://$(docker-machine ip):8080/foo/bar
```
