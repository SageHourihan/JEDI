# JEDI

JEDΙ is a versatile Golang toolkit that empowers you to master the art of Electronic Data Interchange (EDI). Whether you're validating, parsing, translating, or wielding the force of EDI, JEDΙ is your companion in the vast universe of data exchange.

## Features

### 1. Translation

The toolkit allows you to translate EDI files into JSON using the [edx12](https://github.com/arcward/edx12) library.

```bash
myapp translate --file <file_path>
```

Flags

    --file: Path to the EDI file to be translated (required).

### 2. Validation

JEDΙ enables you to validate EDI data, ensuring its compliance with industry standards. The validation feature checks for basic syntax, structural correctness, and other relevant criteria using the [edx12](https://github.com/arcward/edx12) library.

```bash
myapp validate --file <file_path>
```

Flags

    --file: Path to the EDI file to be translated (required).

## Examples

Translating a file

```bash
myapp translate --file path/to/edi/file.edi
```

Output will be written to
``./tmp/output.txt``

Validating a file

```bash
myapp validate --file path/to/edi/file.edi
```

## Contribution

Contributions are welcome! If you have any suggestions, bug reports, or
feature requests, please open an issue or submit a pull request.
