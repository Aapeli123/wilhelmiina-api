FROM golang
WORKDIR /go/src/github.com/Aapeli123/wilhelmiina-api
COPY . .
RUN go get -d -v ./...
RUN go build .
EXPOSE 8080
CMD ["./wilhelmiina-api"]