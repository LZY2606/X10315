package tags

import (
	"io"

	"github.com/osteele/liquid/expressions"
	"github.com/osteele/liquid/render"
)

type fileWriter interface {
	RenderFileTo(io.Writer, string, map[string]any) error
}

func includeTag(source string) (func(io.Writer, render.Context) error, error) {
	expr, err := expressions.Parse(source)
	if err != nil {
		return nil, err
	}

	return func(w io.Writer, ctx render.Context) error {
		_, filename, err := evalTemplatePath(ctx, "include", expr)
		if err != nil {
			return err
		}

		if fw, ok := ctx.(fileWriter); ok {
			return fw.RenderFileTo(w, filename, map[string]any{})
		}

		s, err := ctx.RenderFile(filename, map[string]any{})
		if err != nil {
			return err
		}

		_, err = io.WriteString(w, s)

		return err
	}, nil
}
