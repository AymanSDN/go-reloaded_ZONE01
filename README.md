<h1 align="center" id="title">Go-reloaded</h1>
Certainly! Let's create a Go program that performs the specified modifications on a given input text and writes the modified text to an output file. I'll provide you with a brief description of the project, along with some code snippets to get you started.

---

## Text Modifier Tool

The **Text Modifier Tool** is a command-line utility written in Go that processes input text files and applies various modifications based on predefined rules. It takes two arguments: the input file path and the output file path. The modifications include:

1. **Hexadecimal to Decimal Conversion**:
   - Replace every instance of `(hex)` with the decimal version of the preceding word (assuming the word is a hexadecimal number).
   - Example: `"1E (hex) files were added"` becomes `"30 files were added"`.

2. **Binary to Decimal Conversion**:
   - Replace every instance of `(bin)` with the decimal version of the preceding word (assuming the word is a binary number).
   - Example: `"It has been 10 (bin) years"` becomes `"It has been 2 years"`.

3. **Uppercase Conversion**:
   - Replace every instance of `(up)` with the uppercase version of the preceding word.
   - Example: `"Ready, set, go (up)!"` becomes `"Ready, set, GO!"`.

4. **Lowercase Conversion**:
   - Replace every instance of `(low)` with the lowercase version of the preceding word.
   - Example: `"I should stop SHOUTING (low)"` becomes `"I should stop shouting"`.

5. **Capitalization Conversion**:
   - Replace every instance of `(cap)` with the capitalized version of the preceding word.
   - Example: `"Welcome to the Brooklyn bridge (cap)"` becomes `"Welcome to the Brooklyn Bridge"`.

6. **Custom Word Case Conversion**:
   - If a number appears next to `(low)`, `(up)`, or `(cap)`, apply the specified case conversion to the preceding number of words.
   - Example: `"This is so exciting (up, 2)"` becomes `"This is SO EXCITING"`.

7. **Punctuation Formatting**:
   - Ensure that punctuation marks `.`, `,`, `!`, `?`, `:`, and `;` are correctly placed relative to adjacent words.
   - Handle special cases for groups of punctuation like `...` or `!?`.
   - Example: `"I was sitting over there ,and then BAMM !!"` becomes `"I was sitting over there, and then BAMM!!"`.

8. **Single Quotes Handling**:
   - Place single quotes `'` to the right and left of the word in the middle of them (without spaces).
   - Example: `"I am exactly how they describe me: ' awesome '"` becomes `"I am exactly how they describe me: 'awesome'"`.

9. **Indefinite Article Correction**:
   - Replace every instance of `a` with `an` if the next word begins with a vowel (a, e, i, o, u) or an `h`.
   - Example: `"There it was. A amazing rock!"` becomes `"There it was. An amazing rock!"`.

---

### Usage
Run the program with input and output file paths:
```bash
go run . sample.txt result.txt
```
OR

Compile and Run the Go program:
```bash
go build . && ./go-reloaded sample.txt result.txt
```

---

Feel free to adapt and expand upon this tool as needed. If you have any questions or need further assistance, don't hesitate to ask! 🚀
