FROM golang:1.9.1


WORKDIR /go/src/ClientConsum/appapi
COPY appapi .
COPY mcs ../mcs

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD [ "appapi" ]