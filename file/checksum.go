package file

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"io"
	"os"
)

const bufferSize = 65536

// MD5sumReader returns MD5 checksum of content in reader
func MD5sumReader(reader io.Reader) (string, error) {
	return sumReader(md5.New(), reader)
}

// SHA1sumReader returns MD5 checksum of content in reader
func SHA1sumReader(reader io.Reader) (string, error) {
	return sumReader(sha1.New(), reader)
}

// SHA256sumReader returns SHA256 checksum of content in reader
func SHA256sumReader(reader io.Reader) (string, error) {
	return sumReader(sha256.New(), reader)
}

// sumReader calculates the hash based on a provided hash provider
func sumReader(hashAlgorithm hash.Hash, reader io.Reader) (string, error) {
	buf := make([]byte, bufferSize)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hashAlgorithm.Write(buf[:n])
		case io.EOF:
			return fmt.Sprintf("%x", hashAlgorithm.Sum(nil)), nil
		default:
			return "", err
		}
	}
}

// sum calculates the hash based on a provided hash provider
func sum(hashAlgorithm hash.Hash, filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	return sumReader(hashAlgorithm, bufio.NewReader(file))
}

// MD5sum returns MD5 checksum of filename
func MD5sum(filename string) (string, error) {
	return sum(md5.New(), filename)
}

// SHA1sum returns SHA256 checksum of filename
func SHA1sum(filename string) (string, error) {
	return sum(sha1.New(), filename)
}

// SHA256sum returns SHA256 checksum of filename
func SHA256sum(filename string) (string, error) {
	return sum(sha256.New(), filename)
}
