package testdata

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("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"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
