ARG SWAGGER_2_MARKUP_VERSION=1.3.3

FROM openjdk:8 as swagger2markup
ARG SWAGGER_2_MARKUP_VERSION

WORKDIR /usr/src/

RUN git clone -b v${SWAGGER_2_MARKUP_VERSION} --single-branch --depth 1 https://github.com/Swagger2Markup/swagger2markup-cli.git && \
    cd swagger2markup-cli && \
    ./gradlew assemble



FROM ruby:2.6.3
ARG SWAGGER_2_MARKUP_VERSION

WORKDIR /usr/src

COPY --from=swagger2markup /usr/src/swagger2markup-cli/build/libs/swagger2markup-cli-${SWAGGER_2_MARKUP_VERSION}.jar swagger2markup-cli.jar

RUN apt-get update && apt-get install -y default-jre && \
    gem install asciidoctor-pdf --pre

COPY swagger2markup.properties swagger2pdf.sh ./

ENTRYPOINT [ "/usr/src/swagger2pdf.sh" ]
