FROM golang:1.22-bookworm
WORKDIR /usr/local/bin

COPY . .

EXPOSE 8000
EXPOSE 8126
CMD ["simple-web-server-amd64"]
