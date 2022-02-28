package pulumiutils

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ProjectStackNamer(pulumiCtx *pulumi.Context) func(name string) string {
	ctx := context.Background()
	stack := pulumiCtx.Stack()

	workspace, err := auto.SelectStackInlineSource(ctx, stack, pulumiCtx.Project(), nil)
	if err != nil {
		panic(err)
	}

	acc, err := workspace.Workspace().WhoAmI(ctx)
	if err != nil {
		panic(err)
	}

	return func(name string) string {
		return fmt.Sprintf("%s/%s/%s", acc, name, stack)
	}
}

func ResourceStackNamer(ctx *pulumi.Context) func(name string) string {
	stack := ctx.Stack()

	return func(name string) string {
		return fmt.Sprintf("%s-%s", name, stack)
	}
}
