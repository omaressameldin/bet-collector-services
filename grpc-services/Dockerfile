ARG WORKDIR=/usr/src/app
ARG BUILD_FILE=grpc-service

FROM golang:1.13-alpine
ARG BUILD_FILE
ARG WORKDIR
ARG APP_SRC

WORKDIR $WORKDIR

COPY $APP_SRC/go.mod $APP_SRC/go.sum* ./

RUN apk add git &&\
    apk --update add ca-certificates &&\
    go mod download

COPY ${APP_SRC}/. ./.
RUN CGO_ENABLED=0 GOOS=linux go test ./... &&\
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $BUILD_FILE ./.


FROM bash:4.3.48
ARG BUILD_FILE
ENV BUILD_FILE $BUILD_FILE
ARG WORKDIR
ARG FIREBASE_CONFIG
ENV FIREBASE_CONFIG $FIREBASE_CONFIG
ARG COLLECTION
ENV COLLECTION $COLLECTION
ARG PORT
ENV PORT $PORT


COPY $FIREBASE_CONFIG .
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 ${WORKDIR}/${BUILD_FILE} $BUILD_FILE

RUN addgroup -S appuser && adduser -S appuser -G appuser -u 1000 &&\
    chown -R appuser $BUILD_FILE
USER appuser

EXPOSE $PORT

CMD ./$BUILD_FILE \
    --firebaseConfig=${FIREBASE_CONFIG} \
    --port=${PORT} \
    --collection=${COLLECTION}
