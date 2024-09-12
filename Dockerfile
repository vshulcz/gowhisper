#####################################################
### Copy platform specific binary
FROM bash as copy-binary
ARG TARGETPLATFORM

RUN echo "Target Platform = ${TARGETPLATFORM}"

COPY dist /dist/

RUN if [ "$TARGETPLATFORM" = "linux/amd64" ]; then cp /dist/gowhisper-linux /gowhisper; fi
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then cp /dist/gowhisper-arm /gowhisper; fi
RUN if [ "$TARGETPLATFORM" = "linux/386" ]; then cp /dist/gowhisper-386 /gowhisper; fi
RUN if [ "$TARGETPLATFORM" = "linux/arm/v7" ]; then cp /dist/gowhisper-arm-v7 /gowhisper; fi

RUN chmod +x /gowhisper

#####################################################
### Build Final Image
FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=copy-binary /gowhisper /app/gowhisper

EXPOSE 8080 27017

WORKDIR /app

ENTRYPOINT ["/app/gowhisper"]
