FROM golang:alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags="-s -w \
  -X main.Version=v1.0.0 \
  -X main.Release=production \
  -X main.Sha=$GIT_SHA"

FROM scratch
COPY --from=builder /app/drain /drain

LABEL org.opencontainers.image.title       "gadget"
LABEL org.opencontainers.image.description ""
LABEL org.opencontainers.image.authors     "wilhelm@devilmayco.de"
LABEL org.opencontainers.image.source      "github.com/wilhelm-murdoch/go-gadget"
LABEL org.opencontainers.image.url         "github.com/wilhelm-murdoch/go-gadget"
LABEL org.opencontainers.image.docs        "github.com/wilhelm-murdoch/go-gadget/blob/main/README.md"
LABEL org.opencontainers.image.version     "${GIT_SHA}"

ENTRYPOINT ["gadget", "-h"]