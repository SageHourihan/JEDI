# JEDI

JEDΙ is a versatile Golang toolkit that empowers you to master the art of Electronic Data Interchange (EDI). Whether you're validating, parsing, translating, or wielding the force of EDI, JEDΙ is your companion in the vast universe of data exchange.

## Features
### 1. Translation
The toolkit allows you to translate EDI files by interacting with the Edination API. It supports two main modes:

- **Translate (mode: "t"):** Reads an EDI file, sends it to the Edination API for translation, and writes the translated JSON output to a file.

- **Validate (mode: "v"):** Validates an EDI file using the Edination API and provides feedback on its status

### 2. Validation

JEDΙ enables you to validate EDI data, ensuring its compliance with industry standards. The validation feature checks for basic syntax, structural correctness, and other relevant criteria using the Edination API.

## Examples
Translating a file
``` go run . -t /path/to/edi/file.txt ```

Output will be written to 
``` ./tmp/output.txt ```

Validating a file
``` go run . -v /path/to/edi/file.txt ```
