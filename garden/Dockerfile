FROM golang:1.15

WORKDIR /opt/build
ADD . .
RUN go build -o project .

CMD ["/opt/build/project"]
