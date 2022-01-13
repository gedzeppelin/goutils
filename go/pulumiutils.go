package pulumiutils

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func StackNamer(ctx *pulumi.Context) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s/%s", name, ctx.Stack())
	}
}

func ResNamer(ctx *pulumi.Context) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s-%s", name, ctx.Stack())
	}
}
