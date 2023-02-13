FROM golang:1.18-alpine3.15

ENV MONGO_CONN_STRING=mongodb+srv://doadmin:3Uz59w1m02V76oyk@db-mongodb-blr1-59698-480f7686.mongo.ondigi>

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /zintlr-task-server

EXPOSE 8000

CMD [ "/zintlr-task-server" ]