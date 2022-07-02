FROM golang as build

COPY . /src

WORKDIR /src

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o api_server



FROM scratch as image

LABEL maintainer="ivan.kuchin@gmail.com"
LABEL org.opencontainers.image.source="https://github.com/IvanKuchin/${BUILD_REPO}"

COPY --from=build /src/api_server .
COPY --from=build /src/config ./config

EXPOSE 80

CMD ["/api_server"]
