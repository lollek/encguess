# Encguess

Encguess is an application which attempts to guess the encoding of a given file.
It will be focused on solving practical problems which I have encountered, so it
might be a bit too tailored to me, and not generic enough for general use.

## Current support

Currently, the following encodings are supported from parsing the BOM:
* UTF-1
* UTF-7
* UTF-8
* UTF-16 (Little Endian)
* UTF-16 (Big Endian)
* UTF-32 (Little Endian)
* UTF-32 (Big Endian)
* UTF-EBCDIC
* SCSU
* BOCU-1
* GB-18030

And the following encodings are supported from parsing the rest of the file:
* ASCII
* Printable ASCII

## Requirement
* Go (v1.13+)
* GNU Make (optional)

## Compiling
Either run `make` or `go build .`

## Running

After compiling, the program can be run with ./engcuess FILENAME, where FILENAME
is the name of the file you want to know the encoding of.
