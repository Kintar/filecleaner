package filecleaner

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func HashFile(path string) (string, error) {
	h := md5.New()
	var result string
	f, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		return result, err
	}
	result = hex.EncodeToString(h.Sum(nil))
	return result, nil
}
