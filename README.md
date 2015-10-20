# ts

A simple Go CLI utility to encrypt / decrypt data with TripleSec.


# Usage

It reads the data from the [standard input][stdst], and writes the signed ciphertext to the standard output stream.

```bash
echo "Hello World!" | ./ts -p="TOP SECRET PASSWORD" > encrypted.txt
```


[stdst]: https://en.wikipedia.org/wiki/Standard_streams