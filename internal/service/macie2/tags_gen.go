// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
)

// map[string]*string handling

// Tags returns macie2 service tags.
func Tags(tags tftags.KeyValueTags) map[string]*string {
	return aws.StringMap(tags.Map())
}

// KeyValueTags creates KeyValueTags from macie2 service tags.
func KeyValueTags(ctx context.Context, tags map[string]*string) tftags.KeyValueTags {
	return tftags.New(ctx, tags)
}

// GetTagsIn returns macie2 service tags from Context.
// nil is returned if there are no input tags.
func GetTagsIn(ctx context.Context) map[string]*string {
	if inContext, ok := tftags.FromContext(ctx); ok {
		if tags := Tags(inContext.TagsIn.UnwrapOrDefault()); len(tags) > 0 {
			return tags
		}
	}

	return nil
}

// SetTagsOut sets macie2 service tags in Context.
func SetTagsOut(ctx context.Context, tags map[string]*string) {
	if inContext, ok := tftags.FromContext(ctx); ok {
		inContext.TagsOut = types.Some(KeyValueTags(ctx, tags))
	}
}
