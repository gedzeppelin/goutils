package pulumiutils

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func StackNamer(ctx *pulumi.Context) func(name string) string {
	c := context.Background()
	stack, err := auto.UpsertStackInlineSource(c, ctx.Stack(), ctx.Project(), nil)
	if err != nil {
		panic(err)
	}

	acc, err := stack.Workspace().WhoAmI(c)
	if err != nil {
		panic(err)
	}

	return func(name string) string {
		return fmt.Sprintf("%s/%s/%s", acc, name, ctx.Stack())
	}
}

func ResNamer(ctx *pulumi.Context) func(name string) string {
	return func(name string) string {
		return fmt.Sprintf("%s-%s", name, ctx.Stack())
	}
}
