package flags

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

type BoolFlag struct {
	Key          string
	DefaultValue bool
}

func (f BoolFlag) Value(ctx context.Context, evaluationContext openfeature.EvaluationContext) bool {
	client := openfeature.NewClient("product-catalog")
	value, _ := client.BooleanValue(ctx, f.Key, f.DefaultValue, evaluationContext)
	return value
}

var ProductCatalogFailure = BoolFlag{
	Key:          "productCatalogFailure",
	DefaultValue: false,
}

var ProductCatalogLockContention = BoolFlag{
	Key:          "productCatalogLockContention",
	DefaultValue: false,
}
