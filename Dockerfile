FROM golang:latest

LABEL maintainer="Agus Budianto <agus.kbk@gmail.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build="go build fga" --command="./wait-for-it.sh ${POSTGRES_ADDR}:${POSTGRES_PORT} -- ./fga"
ENTRYPOINT CompileDaemon --build="go build fga" --command="./fga"