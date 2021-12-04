FROM golang:1.16-alpine AS build
ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN ["go", "install", "doximus"]
EXPOSE $PORT
CMD ["doximus", "serve" ,"--port=$PORT"]