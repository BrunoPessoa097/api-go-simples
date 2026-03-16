FROM golang:1.25.7 as stage1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api-golang-simples ./cmd/api


############
FROM scratch
COPY --from=stage1 /app/api-golang-simples /

ENTRYPOINT [ "/api-golang-simples" ]