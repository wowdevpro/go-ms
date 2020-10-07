FROM golang:alpine
WORKDIR /app
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN go get github.com/bitly/go-nsq
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux

RUN go build -o main .

EXPOSE 8082

CMD ["./main"]