package util

import (
	"encoding/base64"
	"math/rand"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/golang-jwt/jwt"
)

func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func URLEscape(s string) string {
	return url.QueryEscape(s)
}

func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func DecodeJwt(encoded string) ([][]byte, error) {
	// split the string into 3 segments and decode
	segments := strings.Split(encoded, ".")
	decoded := [][]byte{}
	for _, segment := range segments {
		bytes, _ := jwt.DecodeSegment(segment)
		decoded = append(decoded, bytes)
	}
	return decoded, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang
// open opens the specified URL in the default browser of the user.
func OpenUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func GetCommit() string {
	bytes, err := exec.Command("git", "rev --parse HEAD").Output()
	if err != nil {
		return ""
	}
	return string(bytes)
}

func Tokenize(s string) map[string]any {
	tokens := make(map[string]any)

	// find token enclosed in curly brackets

	return tokens
}
