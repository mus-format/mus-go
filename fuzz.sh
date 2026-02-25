#!/bin/bash

# Default fuzz time
FUZZTIME=${1:-10s}

echo "Starting fuzz tests with -fuzztime=$FUZZTIME..."

# Find all packages
PACKAGES=$(go list ./...)

for pkg in $PACKAGES; do
  # Get absolute path for the package
  pkg_path=$(go list -f '{{.Dir}}' "$pkg")
  
  # Find all Fuzz functions in the package
  # We look for "func Fuzz" at the beginning of the line in *_test.go files
  fuzz_funcs=$(grep -hPo '^func \K(Fuzz\w+)' "$pkg_path"/*_test.go 2>/dev/null)
  
  if [ -n "$fuzz_funcs" ]; then
    echo "------------------------------------------------------------------------"
    echo "Package: $pkg"
    for func in $fuzz_funcs; do
      echo "Running $func..."
      go test -v -fuzz="^$func$" "$pkg" -fuzztime "$FUZZTIME"
      if [ $? -ne 0 ]; then
        echo "FAIL: $func in $pkg"
        exit 1
      fi
    done
  fi
done

echo "------------------------------------------------------------------------"
echo "All fuzz tests passed!"
