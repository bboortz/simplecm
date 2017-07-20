FROM golang:alpine
RUN apk update && \
	apk add \
		build-base \
		file \
		git && \
	rm -rf /var/cache/apk/*

RUN go get github.com/op/go-logging
RUN go get "github.com/davecgh/go-spew/spew"

RUN mkdir /app 
RUN mkdir /out
ADD . /app
WORKDIR /app 
RUN go fmt .
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scm .
RUN GOOS=linux go build --ldflags "-linkmode external -extldflags -static" -o scm .
#RUN go build -a -installsuffix cgo -o scm .
RUN file scm


CMD ["cp", "/app/scm", "/out"]
