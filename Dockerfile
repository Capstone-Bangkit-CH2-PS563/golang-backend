FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

RUN go mod tidy

COPY . ./

COPY .env ./

RUN go build -o /nutriplant

EXPOSE 8080

CMD [ "/app/nutriplant" ]