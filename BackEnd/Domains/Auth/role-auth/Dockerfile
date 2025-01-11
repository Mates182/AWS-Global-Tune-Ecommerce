FROM rust:1.84 as builder
WORKDIR /usr/src/auth-service
RUN apt-get update && apt-get install -y protobuf-compiler
COPY Cargo.toml Cargo.lock ./
RUN cargo fetch
COPY . .
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /usr/src/auth-service
COPY --from=builder /usr/src/auth-service/target/release/auth-service .
EXPOSE 80
CMD ["./auth-service"]