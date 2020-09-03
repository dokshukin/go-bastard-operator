FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY ./templates .
COPY ./bastard_operator .

EXPOSE 8080/tcp

CMD ["./bastard_operator"]
