FROM golang:alpine as builder

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git

RUN mkdir /crm-integrator
RUN mkdir -p /crm-integrator/proto
WORKDIR /crm-integrator

COPY ./proto/customer.pb.go /crm-integrator/proto
COPY ./cmd/server/main.go /crm-integrator

COPY go.mod .
COPY go.sum .

#RUN go mod download

RUN go build -o crm-integrator .

CMD ./crm-integrator