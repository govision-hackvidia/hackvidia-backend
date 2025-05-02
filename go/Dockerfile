FROM golang:1.24-bookworm AS go_builder
WORKDIR /src
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /govision_main

FROM debian:stable-slim AS go_production
COPY --from=go_builder /govision_main /govision_main
CMD [ "/govision_main" ]
# CMD ["/govision_main"]

