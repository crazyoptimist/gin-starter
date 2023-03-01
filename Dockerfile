##
## STEP 1 - BUILD
##

FROM golang:1.20 as builder

ARG USERNAME=iamuser
ARG USER_UID=1001
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

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

COPY --from=builder /etc/passwd /etc/passwd

USER 1001

EXPOSE 8080

ENTRYPOINT ["/app"]
