# JSON to YAML Converter CLI

This command-line utility converts JSON to YAML and vice versa. It supports specifying input and output files, allowing you to easily convert data between these formats.

## Installation

Clone the repository and build the executable using the following commands:

```bash
git clone <repository-url>
cd <repository-directory>
go build -o converter
```

## Usage

### Convert JSON to YAML

```bash

./converter toyaml -file input.json -output output.yaml


-file: Path to the input JSON file.

-output: (Optional) Path to the output YAML file. If not provided, the result will be printed to the command line.
```

### Convert YAML to JSON

```bash
./converter tojson -file input.yaml -output output.json

-file: Path to the input YAML file.
-output: (Optional) Path to the output JSON file. If not provided, the result will be printed to the command line.
```
## Examples
### Example 1: Convert JSON to YAML and print to the command line

```bash

./converter tojson --file values.yaml --output test.json

```

### Example 2: Convert YAML to JSON and save to a file

```bash

./converter toyaml --file test.json --output test.ya

```
