FROM golang:1.9-alpine as builder

ARG DOCK_PKG_DIR=/go/src/github.com/kyma-project/kyma/components/metadata-service

RUN mkdir -p $DOCK_PKG_DIR

COPY ./ $DOCK_PKG_DIR
WORKDIR $DOCK_PKG_DIR

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o metadataservice ./cmd/metadataservice

FROM alpine
LABEL source=git@github.com:kyma-project/kyma.git

RUN apk update && apk add curl ngrep

COPY --from=builder /go/src/github.com/kyma-project/kyma/components/metadata-service/metadataservice .
COPY ./docs/api/api.yaml .

CMD ["/metadataservice"]
