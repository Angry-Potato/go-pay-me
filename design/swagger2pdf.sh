#!/usr/bin/env bash
set -e

if [ "$#" -lt 1 ]
then
  echo "Usage: swagger2pdf.sh [-s] SWAGGER_FILE <OPTIONAL: -o OUTPUT_DIR -j SWAGGER_2_MARKUP_JAR -c SWAGGER_2_MARKUP_CONFIG>"
  exit 1
fi

dir=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )


while getopts ":o:j:c:s:" opt; do
  case $opt in
    s) cli_s="$OPTARG"
    ;;
    o) cli_o="$OPTARG"
    ;;
    j) cli_j="$OPTARG"
    ;;
    c) cli_c="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    ;;
  esac
done

swaggerfile=${cli_s:-$1}
outdir=${cli_o:-`dirname ${swaggerfile}`}
swagger2markupjar=${cli_j:-${dir}/swagger2markup-cli.jar}
swagger2markupconfig=${cli_c:-${dir}/swagger2markup.properties}
swagger2markupout=${outdir}/index

java -jar ${swagger2markupjar} convert \
    -i ${swaggerfile} \
    -f ${swagger2markupout} \
    -c ${swagger2markupconfig}

asciidoctor-pdf ${swagger2markupout}.adoc \
    -a toc=left \
    -a toclevels=3 \
    -a sectanchors \
    --out-file=${outdir}/index.pdf