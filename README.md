# ts

A simple Go CLI utility to encrypt / decrypt data with [TripleSec][ts].

---

[TripleSec][ts] is a simple, triple-paranoid, symmetric encryption library created by [Keybase][keybase].

It encrypts data with Salsa 20, AES, and Twofish, so that a compromise of one or two of the ciphers will not expose the secret.


## Usage

It reads the data from the [standard input][stdst], and writes the signed ciphertext to the [standard output][stdst] stream.

### Encrypting

```bash
echo "Hello World!" | ./ts -p="TOP SECRET PASSWORD" > encrypted.txt
[TripleSec] Encrypting...
```

### Decrypting

```bash
cat encrypted.txt | ./ts -decrypt -p="TOP SECRET PASSWORD"
[TripleSec] Decrypting...
Hello World
```


[ts]: https://keybase.io/triplesec/
[keybase]: https://keybase.io
[stdst]: https://en.wikipedia.org/wiki/Standard_streams