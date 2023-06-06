FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "be-api"
RUN go build -o immersive-dashboard-api

# run executable file
CMD ["./immersive-dashboard-api"]