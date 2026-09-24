package tags

import (
	"io"

	"github.com/osteele/liquid/render"
)

// fileWriter is implemented by render contexts that can render a template
// file directly to a writer, within the caller's lexical scope.
type fileWriter interface {
	RenderFileTo(io.Writer, string, map[string]any) error
}

// isolatedFileRenderer is implemented by render contexts that can render a
// template file without inheriting the caller's lexical scope.
type isolatedFileRenderer interface {
	RenderFileIsolated(string, map[string]any) (string, error)
}

// isolatedFileWriter is the writer-based variant of isolatedFileRenderer.
type isolatedFileWriter interface {
	RenderFileIsolatedTo(io.Writer, string, map[string]any) error
}

// writeIncludedFile renders a template file into w, within the caller's
// lexical scope. It is the shared file-rendering boundary for {% include %}.
func writeIncludedFile(w io.Writer, ctx render.Context, filename string, bindings map[string]any) error {
	if fw, ok := ctx.(fileWriter); ok {
		return fw.RenderFileTo(w, filename, bindings)
	}

	s, err := ctx.RenderFile(filename, bindings)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, s)

	return err
}

// writeIsolatedFile renders a template file into w, without inheriting the
// caller's lexical scope. It is the shared file-rendering boundary for the
// {% render %} tag's "with" and "for" forms.
func writeIsolatedFile(w io.Writer, fileRenderer isolatedFileRenderer, filename string, scope map[string]any) error {
	if fw, ok := fileRenderer.(isolatedFileWriter); ok {
		return fw.RenderFileIsolatedTo(w, filename, scope)
	}

	s, err := fileRenderer.RenderFileIsolated(filename, scope)
	if err != nil {
		return err
	}

	_, err = io.WriteString(w, s)

	return err
}
