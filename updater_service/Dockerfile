FROM golang

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/update

EXPOSE 80
RUN chmod +x /bin/update
ENTRYPOINT update