package signature

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strings"
)

type HmacSha1Sign struct {
	rawStringGenerater RawStringGenerater
}

func (s HmacSha1Sign) Signature(rawParams []string, key string) string {
	if s.rawStringGenerater == nil {
		panic("need new a RawStringGenerater")
	}
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(s.rawStringGenerater.RawString(rawParams)))
	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

func (s HmacSha1Sign) GetRawStringGenerater() (RawStringGenerater, error) {
	if s.rawStringGenerater == nil {
		return nil, errors.New("do not have any RawStringGenerater")
	}
	return s.rawStringGenerater, nil
}

func (s HmacSha1Sign) SetRawStringGenerater(gen RawStringGenerater) error {
	s.rawStringGenerater = gen
	return nil
}

func (s HmacSha1Sign) Check(rawParams []string, key, encodedString string) bool {
	return strings.Compare(s.Signature(rawParams, key), encodedString) == 0
}
