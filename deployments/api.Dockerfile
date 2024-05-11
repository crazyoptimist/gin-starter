##
## STAGE 1 - BUILD
##

FROM golang:1.22 as builder

# Add a non-root user (for prod image)
ARG USERNAME=iamuser
ARG USER_UID=1001
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME

# Required for performing outbound HTTP requests (for prod image)
RUN apt update && apt install -y ca-certificates

ENV APP_HOME /source

WORKDIR "$APP_HOME"

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/cli ./cmd/cli

##
## STAGE 2 - PRODUCTION
##

FROM scratch

COPY --from=builder /source/bin/server /server
COPY --from=builder /source/bin/cli /cli

ENV GIN_MODE=release
ENV TWELVE_FACTOR_MODE=true

# Copy the non-root user and ca-certificates
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER 1001

EXPOSE 8080

ENTRYPOINT ["/server"]
