#
# Multi-stage build for Angular app and Go app
#


# Build Angular app in a node container
FROM node:12-alpine as ngbuilder
RUN npm install -g @angular/cli
RUN mkdir /src
WORKDIR /src
COPY snaily-web /src/snaily-web
WORKDIR /src/snaily-web
#RUN chmod -R 777 ./
RUN npm i
RUN ng build

# Build Golang app in Go container
FROM golang:1.12.4-alpine3.9 as builder
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Final app
FROM scratch as final
COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /app
USER nobody:nobody
COPY --chown=nobody:nobody --from=ngbuilder /src/snaily-web/dist/snaily-web /snaily-web
ENTRYPOINT ["/app", "-container=true", "-pgport=5432", "-pghost=spg", "-nsqd=nsq_nsqd:4150"]