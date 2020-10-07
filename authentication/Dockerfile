FROM golang:alpine
WORKDIR /app
COPY . .

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]