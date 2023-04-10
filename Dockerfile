FROM golang:1.20-alpine as build-stage

ENV CGO_ENABLED 0
ENV PROJECT_PACKAGE profit-earnings
ENV OBJ_NAME profit-earnings

COPY . /go/src/${PROJECT_PACKAGE}/

RUN cd /go/src/${PROJECT_PACKAGE} && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -o ${OBJ_NAME}

# --------------------------------------------------------------------------------
FROM alpine:3.17

RUN apk add --no-cache bash

ENV PROJECT_PACKAGE profit-earnings
ENV OBJ_NAME profit-earnings

RUN adduser -D ${OBJ_NAME}
USER ${OBJ_NAME}

COPY --from=build-stage /go/src/${PROJECT_PACKAGE}/${OBJ_NAME} /usr/local/bin/${OBJ_NAME}
COPY --from=build-stage /go/src/${PROJECT_PACKAGE}/test_cases /usr/local/bin/test_cases
COPY --from=build-stage /go/src/${PROJECT_PACKAGE}/run_test_cases.sh /usr/local/bin
WORKDIR /usr/local/bin/
CMD bash run_test_cases.sh