package signature

type Signer interface {
	GetRawStringGenerater() (RawStringGenerater, error)
	SetRawStringGenerater(gen RawStringGenerater) error
	Signature(rawParams []string, key string) string
	Check(rawParams []string, key, encodedString string) bool
}
