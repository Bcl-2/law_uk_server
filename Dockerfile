FROM golang:1.21.5

WORKDIR /app_server
COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLEd=0 GOOS=linix go build -o /docker-gs-ping

CMD ["/docker-gs-ping"]