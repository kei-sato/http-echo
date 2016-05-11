# getting started

```
docker pull keisato/http-echo && docker tag keisato/http-echo http-echo
docker run --rm -it --name http-echo -p 8080:80 http-echo
```

# output

```
$ curl -iXPOST -d hoge=111 http://$(docker-machine ip):8080/foo/bar
HTTP/1.1 200 OK
Date: Wed, 11 May 2016 02:24:52 GMT
Content-Length: 183
Content-Type: text/plain; charset=utf-8

POST /foo/bar HTTP/1.1
Host: 192.168.99.100:8080
Content-Length: 8
Accept: */*
Content-Length: 8
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.43.0

hoge=111
```

# development

## run 

```
PORT=8080 go run main.go
curl -i http://localhost:8080/foo/bar
```

## create image

```
docker build http-echo .
docker run --rm -p 8080:80 --name http-echo http-echo
curl -i http://$(docker-machine ip):8080/foo/bar
```
