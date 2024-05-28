#!/usr/bin/env bash

org=cycloid-sandbox


export CY_CREATE_ENV_VARS='{"viaEnv": "ok", "file4": "overriden", "subMap": {"two": 2}}'
echo -e '{\n"file1": "titi",\n"toto": "toto"\n}{"file4": 3, "subMap": {"one": 1}}' | go run . projects create-stackforms-env \
  --verbosity debug \
  --org $org \
  --project "test-env" \
  --env "test" \
  --vars '{"toto": "vars1"}' \
  --var-file '-' \
  --var-file <(echo '{"file2": "tata"}')
