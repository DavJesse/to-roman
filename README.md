# to-roman
This program converts decimal numbers parsed as input to their roman representation

## Usage
To use this program execute a command in your terminal with the following format:
```
bash
$ go run . 1234
```
Where `1234` is the the number you wish to convert to Roman numerals. Adding more arguments (or less) will print out an error and exit the program.

### Rules
The number you type in a input should adhere to the following rules:
- Not contain non-numeric characters like `','`, `'.'`, `'+'`, `'-'`, etc.
- Be greater than '0'
- Be less than '4000'

### Output
Apart from displaying the Roman representation of the integer parsed as input, the the program prints out an overview of how the result was calculated.

Refer to the box below:
```
bash
MMIX
M+M+(I-X)
```
### Requirements
`to-roman` is a stand-alone program written on go. To run it, you will need to have Golang insalled in your computer
