FROM golang:1.16-alpine3.12 as lisn_builder
RUN mkdir /lisn
WORKDIR /lisn
RUN apk --no-cache add ca-certificates
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o bin
EXPOSE 8000 80
CMD [ "/lisn/bin", "--config", "./configs/production.yml" ]