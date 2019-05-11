# Swagger2PDF

Swagger2PDF is a simple shell script that will convert a given swagger / openapi2.0 JSON file to a PDF file. Heavily inspired by [openapi-asciidoctor](https://github.com/Ajaxy/openapi-asciidoctor).

## Requirements

To run Swagger2PDF, you will need:

- [OpenJDK 8](https://openjdk.java.net/install/)
- [Swagger2Markup-CLI](https://github.com/Swagger2Markup/swagger2markup-cli)
- [RubyGem asciidoctor-pdf](https://asciidoctor.org/docs/asciidoctor-pdf/#install-the-published-gem)

**OR** you can use Docker.

## Running

To run the script, provide it with:

- SWAGGER_FILE (**required** flag: -s) = a path to the `swagger.json` you wish to convert
- OUTPUT_DIR (**optional** flag: -o) = the directory you wish the generated files to appear under. Defaults to the directory of the given swagger file.
- SWAGGER_2_MARKUP_JAR (**optional** flag: -j) = the swagger2markup-cli.jar you wish to use. Defaults to `/swagger2markup-cli.jar` next to where the `swagger2pdf.sh` script is located.
- SWAGGER_2_MARKUP_CONFIG (**optional** flag: -c) = the swagger2markup config properties you wish to use. Defaults to `/swagger2markup.properties` next to where the `swagger2pdf.sh` script is located.

e.g.:

    $ ./swagger2pdf.sh -s swagger.json -c swagger2markup.properties

**OR** you can use the given Dockerfile. Build it like so:

    $ docker build --pull --tag swagger2pdf:latest .

And then run it like so:

    (assuming you have a swagger.json in your current directory)
    $ docker run -v ${PWD}:/tmp/s2pdf swagger2pdf:latest /tmp/s2pdf/swagger.json /tmp/s2pdf

After running, the generated files should appear under the specified directory.
