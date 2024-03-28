# JEDI

JEDΙ is a versatile Golang toolkit that empowers you to master the art of Electronic Data Interchange (EDI). Whether you're validating, parsing, translating, or wielding the force of EDI, JEDΙ is your companion in the vast universe of data exchange.

## Features
### 1. Translation
The toolkit allows you to translate EDI files into JSON using the [edx12](https://github.com/arcward/edx12) library.

- **Translate (mode: "t"):** Takes an EDI file and writes the translated JSON output to a file.

### 2. Validation

JEDΙ enables you to validate EDI data, ensuring its compliance with industry standards. The validation feature checks for basic syntax, structural correctness, and other relevant criteria using the [edx12](https://github.com/arcward/edx12) library.
- **Validate (mode: "v"):** Validates an EDI file and provides feedback on its status

## Examples
Translating a file
``` go run . -t /path/to/edi/file.txt ```

Output will be written to 
``` ./tmp/output.txt ```

Validating a file
``` go run . -v /path/to/edi/file.txt ```
