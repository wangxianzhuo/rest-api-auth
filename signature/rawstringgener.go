package signature

type RawStringGenerater interface {
	RawString(params []string) string
}
