# Build stage: full Go toolchain, discarded after compiling.
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/server ./cmd

# Runtime stage: just the static binary + the frontend, running as non-root.
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/server ./server
# main.go serves http.Dir("static") relative to the working directory.
COPY static ./static
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/app/server"]
