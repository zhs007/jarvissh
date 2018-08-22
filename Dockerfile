
FROM golang:1.10 as builder

MAINTAINER zerro "zerrozhao@gmail.com"

WORKDIR $GOPATH/src/github.com/zhs007/jarvissh

COPY . $GOPATH/src/github.com/zhs007/jarvissh

RUN go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jarvissh . \
    && mkdir /home/jarvissh \
    && cp jarvissh /home/jarvissh

FROM scratch
WORKDIR /home/jarvissh
COPY --from=builder /home/jarvissh /home/jarvissh
CMD ["./jarvissh"]
