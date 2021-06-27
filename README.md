# External sorting

An example application to sort huge string files which do not fit to the RAM.

## How to use

Build the binaries:

```bash
make build
```

Use `generator` to create and populate the input file

```bash
# Print help
bin/generator -h

# Run with default parameters
bin/generator

# Customize the generator
bin/generator -max_len=100 -lines=500000 -out=huge.txt
```

Use `sorter` to sort the input file

```bash
# Print help
bin/sorter -h

# Run the sorter
bin/sorter -in huge.txt -out huge_sorted.txt
```