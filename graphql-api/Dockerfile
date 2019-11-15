FROM gradle
ARG WORKDIR=/usr/src/app/api
ARG PROTO_DIR

WORKDIR $WORKDIR

COPY ./graphql-api/api/. ./.
COPY ./graphql-api/grpc/. ./grpc/.
COPY $PROTO_DIR/**/* ./grpc/src/main/proto/.

CMD  gradle \
      && gradle wrapper \
      && ./gradlew run
