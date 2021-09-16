FROM golang:1.17-alpine3.14
RUN mkdir /lisn
WORKDIR /lisn
RUN apk --no-cache add ca-certificates
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o bin
EXPOSE 8000 80
CMD [ "/lisn/bin", "--config", "./configs/production.yml" ]