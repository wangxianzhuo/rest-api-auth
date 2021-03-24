package signature

import "strings"

type CommonRawStringGenerate struct {
}

func (g CommonRawStringGenerate) RawString(params []string) string { return strings.Join(params, "") }
