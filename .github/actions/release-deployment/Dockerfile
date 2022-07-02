FROM golang as build

COPY . /src

WORKDIR /src

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o api_server



FROM scratch as image

COPY --from=build /src/api_server .
COPY --from=build /src/config ./config

EXPOSE 80

CMD ["/api_server"]
