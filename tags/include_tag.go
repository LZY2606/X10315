package tags

import (
	"io"

	"github.com/osteele/liquid/expressions"
	"github.com/osteele/liquid/render"
)

func includeTag(source string) (func(io.Writer, render.Context) error, error) {
	expr, err := expressions.Parse(source)
	if err != nil {
		return nil, err
	}

	return func(w io.Writer, ctx render.Context) error {
		value, err := ctx.Evaluate(expr)
		if err != nil {
			return err
		}

		rel, ok := value.(string)
		if !ok {
			return ctx.Errorf("include requires a string argument; got %v", value)
		}

		filename, err := resolveTemplatePath(ctx.SourceFile(), rel)
		if err != nil {
			return ctx.WrapError(err)
		}

		return writeIncludedFile(w, ctx, filename, map[string]any{})
	}, nil
}
