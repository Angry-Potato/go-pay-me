# :money_with_wings: Go-Pay-Me :money_with_wings: ![Travis (.org)](https://img.shields.io/travis/Angry-Potato/go-pay-me.svg?style=flat-square)

Simple payments resource API complete with design, implementation, automation, live docs, live site, and development journal.

# Table of Contents :books:

- [Final Outputs](#final-outputs-potato)
- [Design](#design-computer)
  - [Docs](#docs-scroll)
- [Implementation](#implementation-weight_lifting_man)
  - [Building](#building-building_construction)
  - [Testing](#testing-performing_arts)

## Final Outputs :potato:

The final outputs, as specified by the brief, are:

- [Design PDF](design/index.pdf)
- [Implementation](implementation)

### Bonus :gem:

Additional deliverables completed:

- [Design docs hosted as Github Page](https://angry-potato.github.io/go-pay-me/)
- [Automation](.travis.yml)
- [Dev journal](JOURNAL.md)
- [Implementation Deployed to Heroku](https://go-pay-me.herokuapp.com/payments)
- [Swagger2PDF](https://github.com/Angry-Potato/swagger2pdf)

## Design :computer:

The design directory contains [the API design](design/payments.swagger.json) in `swagger.json` format, as well as the resulting [PDF](design/index.pdf) as specified in the brief, a copy of the original payments payload example, some swagger2pdf config, and a Makefile to encapsulate the PDF generation.

The PDF, asciidoc, and HTML docs are all autogenerated from the `payments.swagger.json`. This is done using an `npm script` called `regen-docs` in the root of the repository.

### Docs :scroll:

As mentioned above, the docs under the [docs](docs) directory are autogenerated from the [payments.swagger.json](design/payments.swagger.json) using an `npm script`.

To run the generation yourself, clone the repo, cd to the root of it, and install the `node_modules` via:

    yarn # or `npm i` depending on your tool preference

Then run the script using:

    yarn regen-docs # or `npm run regen-docs`

The [docs](docs) dir now contains up-to-date HTML docs based on the [payments.swagger.json](design/payments.swagger.json), and the design dir contains up-to-date PDF and asciidocs.

## Implementation :weight_lifting_man:

The implementation directory contains the API implementation in golang. I used [dep](https://golang.github.io/dep/) for dependency management because, in my experience, it is the most reliable dependency management system for golang and I didn't want to waste time fiddling with dependency versions.

The golang app is built, tested, and deployed using a multistage Dockerfile. The production stage of the image was initially `FROM scratch` to be as small as possible but some minimal tooling was required in the image to have the container reliably testable, `prod-test` and `prod` stages are not a good idea because we should test exactly what runs in production, so the production image is based from alpine instead.

There are two docker-compose files:

- [docker-compose.yml](implementation/docker-compose.yml) - used as a way of running the app locally, consists of the app, and a postgres db.
- [docker-compose.test.yml](implementation/docker-compose.test.yml) - used as a way of full-stack testing the app locally, consists of the app, a tester app, and a postgres db.

### Building :building_construction:

The following instructions assume you are in the [implementation](implementation) directory.

To build the app, generate the `app` binary by running the make command:

    make build

To build the docker image containing the production-ready app and its' test stage, run the make command:

    make build-docker-image

### Testing :performing_arts:

The following instructions assume you are in the [implementation](implementation) directory.

To test the app, execute the unit test suite by running the make command:

    make test

To execute the full-stack tests, run the make command:

    make docker-compose-test

Success or failure can be seen in the logs output, and in the exit code returned by the command.
