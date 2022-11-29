FROM golang:1.17-alpine
RUN mkdir /securities-search
WORKDIR /securities-search

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

EXPOSE 8080

RUN GOOS=linux GOARCH=amd64 go build -o /securities-search/exec ./cmd/securities-search

CMD ["/securities-search/exec"]
