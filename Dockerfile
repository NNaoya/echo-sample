FROM golang:1.15.3

RUN go get github.com/labstack/echo/...
RUN go get github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/echo-swagger

WORKDIR /app
ADD . /app

CMD ["go", "run", "main.go"]