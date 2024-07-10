package gen

import (
	"github.com/cat3306/ginctl/model/sql/template"
	"github.com/cat3306/ginctl/util"
	"github.com/cat3306/ginctl/util/pathx"
)

func genImports(table Table, withCache, timeImport bool) (string, error) {
	if withCache {
		text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
		if err != nil {
			return "", err
		}

		buffer, err := util.With("import").Parse(text).Execute(map[string]any{
			"time":       timeImport,
			"containsPQ": table.ContainsPQ,
			"data":       table,
		})
		if err != nil {
			return "", err
		}

		return buffer.String(), nil
	}

	text, err := pathx.LoadTemplate(category, importsWithNoCacheTemplateFile, template.ImportsNoCache)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]any{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func gormGenImports(timeImport bool, sqlNull bool) (string, error) {
	text, err := pathx.LoadTemplate(category, importGormTemplateFile, template.GormImports)
	if err != nil {
		return "", err
	}

	buffer, err := util.With("import").Parse(text).Execute(map[string]any{
		"time":    timeImport,
		"sqlNull": sqlNull,
	})
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
