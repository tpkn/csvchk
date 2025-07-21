Fast and simple csv file checker. Just pipe csv content into it.


## Usage

```text
csvchk [ -c ] [ -q ] < <file.csv>
```


## Options

```text
-c             Collect and print a complete list of problems with a csv file (default: only first error found)
-q             Just exit with exit code 1
-d             Fields separator (default: comma)
-h, --help     Help
-v, --version  Version
```


## Examples
```shell
# List all errors
csvchk -c < file.csv

# Check gzipped csv file
gunzip -c "file.csv.gz" | csvchk && echo "ok" || echo "(!) error"
```




