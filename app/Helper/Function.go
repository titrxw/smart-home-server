package helper

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	mathRand "math/rand"
	"strings"
	"time"

	"github.com/gookit/goutil/strutil"
	uuid "github.com/satori/go.uuid"
)

func UUid() string {
	return uuid.NewV4().String()
}

func RandomStr(len int) string {
	return strutil.RandomCharsV3(len)
}

func RandomNumber(len int) string {
	AlphaNum := "0123456789"

	cs := make([]byte, len)
	for i := 0; i < len; i++ {
		mathRand.Seed(time.Now().UnixNano())
		idx := mathRand.Intn(10)
		cs[i] = AlphaNum[idx]
	}

	return string(cs)
}

func Sha1(str string) string {
	hash := sha1.New()
	_, _ = io.WriteString(hash, str)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func JsonEncode(data interface{}) (string, error) {
	encodeData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(encodeData[:]), nil
}

func JsonDecode(data string, obj interface{}) error {
	err := json.Unmarshal([]byte(data), &obj)
	if err != nil {
		return err
	}

	return nil
}

func GetSensitiveWord(word string, sensitiveWordMap []string) []string {
	var words []string

	word = strings.TrimSpace(word)
	if word == "" {
		return words
	}

	for _, value := range sensitiveWordMap {
		if strings.Index(word, value) != -1 {
			words = append(words, value)
		}
	}

	return words
}
