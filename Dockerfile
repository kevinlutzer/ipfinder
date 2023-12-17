FROM golang:alpine

# Add Server and Set Working Directory 
ADD . /app/server
WORKDIR /app/server

# Build in Container
RUN go build -o main main.go 
EXPOSE 80

ENTRYPOINT [ "/app/server/main" ]