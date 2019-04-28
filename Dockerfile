
FROM golang:1.12 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR /src/jarvissh

COPY ./go.* /src/jarvissh/

RUN go mod download

COPY . /src/jarvissh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvissh . \
    && mkdir /app \
    && mkdir /app/jarvissh \
    && mkdir /app/jarvissh/dat \
    && mkdir /app/jarvissh/logs \
    && mkdir /app/jarvissh/cfg \
    && cp -r www /app/jarvissh/www \
    && cp ./jarvissh /app/jarvissh/ \
    && cp cfg/config.yaml.default /app/jarvissh/cfg/config.yaml

FROM alpine
WORKDIR /app/jarvissh
COPY --from=builder /app/jarvissh /app/jarvissh
CMD ["./jarvissh", "start"]
