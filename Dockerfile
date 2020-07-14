from golang:latest

WORKDIR /build
copy . .
RUN go get -u ./...
RUN go build -o built .

ENV MONKEBASE_CONNECTION ""
ENTRYPOINT ./built
