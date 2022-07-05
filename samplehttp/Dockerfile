FROM golang:alpine 
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
#final stage
#FROM alpine:latest

#CMD mkdir app
#COPY --from=builder /go/bin/main /app

ENTRYPOINT /go/bin/demo
LABEL Name=samplehttp Version=0.0.1
LABEL maintainer="readyGo team. JitenP@Outlook.Com"
EXPOSE 50059