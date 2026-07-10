package main

const Help = `CSV Check (v%v) | https://tpkn.me

A simple and fast CLI for checking CSV files for errors.

Usage:
  csvchk [ -c ] [ -q ] < <file.csv>

Options:
  -c             Collect all csv errors and output the list at the end
  -q             Silently terminate with exit(1) upon the first error encountered in the CSV
  -d             Fields separator (default: comma)
  -h, --help     Help
  -v, --version  Version

Examples:
  # List all errors
  csvchk -c < file.csv

  # Check gzipped csv file
  gunzip -c "file.csv.gz" | csvchk && echo "ok" || echo "(!) error"
`
