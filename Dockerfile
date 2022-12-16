FROM golang:1.17.1

# GO module configuration
ENV GO111MODULE=on

# Move to project root
WORKDIR /go/src/api

# Install rerun tool
RUN go get -u github.com/ivpusic/rerun

# Copy go mode files
COPY go.mod ./
COPY go.sum ./

# Other non-vendored files
COPY main.go ./
COPY internal internal

EXPOSE 3000


# Install server application
RUN ["go", "install", "."]

CMD [ "pizza-app", "api" ]
