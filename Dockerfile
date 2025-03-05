# Build stage
FROM golang:1.23.1-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o ./bin/signer.app ./pkg/program.go

#exec stage
FROM ubuntu:24.04
WORKDIR /app

# Instalar LibreOffice
RUN apt-get update && apt-get -y upgrade && apt-get install -y libreoffice fonts-dejavu fonts-liberation ttf-mscorefonts-installer

#copy binary file
COPY --from=build /app/bin/signer.app /app/bin/signer.app

# copy requirement directories
COPY --from=build /app/assets /app/assets
COPY --from=build /app/store /app/store

CMD ["/app/bin/signer.app"]