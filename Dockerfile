FROM golang:1.18-alpine AS build

WORKDIR /build
# download the required Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#COPY *.go ./
COPY . ./
RUN ls
RUN go build -o deals-backend

FROM alpine
RUN adduser -S -D -H -h /app appuser

USER appuser
COPY --from=build /build/deals-backend /app/
COPY docs/ /app/docs
COPY .env /app
WORKDIR /app

CMD ["./deals-backend"]