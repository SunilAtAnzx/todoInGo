#!/bin/bash

function collect_coverage() {

COVERAGE_DATA_PATH=coverage.integration.out
COVERAGE_COLLECTOR_URL=http://collector-service/api/upload

if [ ! -f "$COVERAGE_DATA_PATH" ]; then
  echo "[trap:EXIT]- coverage data (path: $COVERAGE_DATA_PATH) does not exist."
  exit
fi

if [ -z "$COVERAGE_COLLECTOR_URL" ]; then
  echo "[trap:EXIT]- collector URL not present, dumping coverage data..."
  echo "--- BEGIN COVERAGE ---"
  cat $COVERAGE_DATA_PATH
  echo "--- END COVERAGE ---"
  exit
fi

echo "[trap:EXIT]- pushing coverage data (url: $COVERAGE_COLLECTOR_URL) to collector."

 curl --location $COVERAGE_COLLECTOR_URL \
--header 'Content-Type: multipart/form-data' \
--form 'coverageReport='@$COVERAGE_DATA_PATH

}

function runInstrumentedBinary() {
  trap "collect_coverage" EXIT
  ./bin/todoInGo.test -test.coverprofile coverage.integration.out
}

runInstrumentedBinary