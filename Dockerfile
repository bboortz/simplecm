FROM golang:alpine
RUN apk update && \
	apk add git && \
	rm -rf /var/cache/apk/*

RUN go get github.com/op/go-logging

RUN mkdir /app 
RUN mkdir /out
ADD . /app
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scm .


CMD ["cp", "/app/scm", "/out"]
