#!/usr/bin/env bash
# Check every episode module with the Go toolchain in $GO (default: go on PATH).
#  - expected.txt: the build must FAIL, and its output must contain every line of expected.txt
#  - output.txt:   `go run .` must print exactly output.txt
#  - otherwise:    gofmt-clean, `go vet` and `go build` must pass
set -uo pipefail
cd "$(dirname "$0")"
GO=${GO:-go}
GOFMT="$("$GO" env GOROOT)/bin/gofmt"
"$GO" version
fail=0
for dir in l*/; do
  dir=${dir%/}
  [ -f "$dir/go.mod" ] || continue
  if [ -f "$dir/expected.txt" ]; then
    out=$(cd "$dir" && "$GO" build -o /dev/null . 2>&1); status=$?
    missing=$(grep -v '^$' "$dir/expected.txt" | while IFS= read -r want; do grep -qF -- "$want" <<<"$out" || echo "$want"; done)
    if [ $status -ne 0 ] && [ -z "$missing" ]; then echo "ok    $dir (fails as expected)"
    else echo "FAIL  $dir (exit $status; missing: ${missing:-none})"; echo "$out"; fail=1; fi
    continue
  fi
  unformatted=$("$GOFMT" -l "$dir")
  out=$(cd "$dir" && "$GO" vet . 2>&1 && "$GO" build -o /dev/null . 2>&1); status=$?
  if [ -n "$unformatted" ] || [ $status -ne 0 ]; then
    echo "FAIL  $dir"; [ -n "$unformatted" ] && echo "not gofmt-ed: $unformatted"; echo "$out"; fail=1; continue
  fi
  if [ -f "$dir/output.txt" ]; then
    got=$(cd "$dir" && "$GO" run . 2>&1)
    if [ "$got" = "$(cat "$dir/output.txt")" ]; then echo "ok    $dir (builds, output matches)"
    else echo "FAIL  $dir (output differs)"; diff <(echo "$got") "$dir/output.txt"; fail=1; fi
  else
    echo "ok    $dir (builds)"
  fi
done
exit $fail
