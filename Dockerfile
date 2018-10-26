
FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/jarvissh

COPY ./Gopkg.* $GOPATH/src/github.com/zhs007/jarvissh/

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure -vendor-only -v

COPY . $GOPATH/src/github.com/zhs007/jarvissh

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvissh . \
    && mkdir /home/jarvissh \
    && mkdir /home/jarvissh/dat \
    && mkdir /home/jarvissh/logs \
    && mkdir /home/jarvissh/cfg \
    && cp jarvissh /home/jarvissh \
    && cp cfg/config.yaml.default /home/jarvissh/cfg/config.yaml

FROM scratch
WORKDIR /home/jarvissh
COPY --from=builder /home/jarvissh /home/jarvissh
CMD ["./jarvissh"]
