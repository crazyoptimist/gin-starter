# syntax=docker/dockerfile:1

##
## STEP 1 - BUILD
##

FROM golang:1.20 as builder

ENV APP_HOME /source

WORKDIR "$APP_HOME"

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/app .

##
## STEP 2 - DEPLOY
##

FROM scratch

COPY --from=builder /source/bin/app /app

COPY .env /.env

EXPOSE 8080

ENTRYPOINT ["/app"]
