FROM openjdk:8
ENV SWAGGER_2_MARKUP_VERSION=1.3.3

WORKDIR /usr/src/

RUN git clone -b v${SWAGGER_2_MARKUP_VERSION} --single-branch --depth 1 https://github.com/Swagger2Markup/swagger2markup-cli.git && \
    cd swagger2markup-cli && \
    ./gradlew assemble