package tags

import (
	"fmt"
	"path/filepath"

	"github.com/osteele/liquid/expressions"
	"github.com/osteele/liquid/render"
)

// evalTemplatePath evaluates a template name expression and resolves it to a
// path relative to the current source file. It is the single entry shared by
// the {% include %} and {% render %} tags, so that both AST nodes judge their
// target the same way before diverging on scope (inherited vs. isolated).
// It returns both the evaluated name and the resolved path.
func evalTemplatePath(ctx render.Context, tagName string, expr expressions.Expression) (string, string, error) {
	value, err := ctx.Evaluate(expr)
	if err != nil {
		return "", "", err
	}

	rel, ok := value.(string)
	if !ok {
		return "", "", ctx.Errorf("%s requires a string template name; got %v", tagName, value)
	}

	filename, err := resolveTemplatePath(ctx.SourceFile(), rel)
	if err != nil {
		return "", "", ctx.WrapError(err)
	}

	return rel, filename, nil
}

func resolveTemplatePath(sourceFile, templateName string) (string, error) {
	if templateName == "" {
		return "", fmt.Errorf("template path is empty")
	}
	if filepath.IsAbs(templateName) || filepath.VolumeName(templateName) != "" {
		return "", fmt.Errorf("template path %q must be relative", templateName)
	}

	clean := filepath.Clean(templateName)
	if !filepath.IsLocal(clean) {
		return "", fmt.Errorf("template path %q escapes its source directory", templateName)
	}

	return filepath.Join(filepath.Dir(sourceFile), clean), nil
}
