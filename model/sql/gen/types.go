package gen

import (
	"github.com/cat3306/ginctl/model/sql/template"
	"github.com/cat3306/ginctl/util"
	"github.com/cat3306/ginctl/util/pathx"
	"github.com/cat3306/ginctl/util/stringx"
)

func genTypes(table Table, methods string, withCache bool) (string, error) {
	fields := table.Fields
	fieldsString, err := genFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesTemplateFile, template.Types)
	if err != nil {
		return "", err
	}

	output, err := util.With("types").
		Parse(text).
		Execute(map[string]any{
			"withCache":             withCache,
			"method":                methods,
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"fields":                fieldsString,
			"data":                  table,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func genGormTypes(table Table) (string, error) {
	fields := table.Fields
	fieldsString, err := genGormFields(table, fields)
	if err != nil {
		return "", err
	}

	text, err := pathx.LoadTemplate(category, typesGormTemplateFile, template.GormTypes)
	if err != nil {
		return "", err
	}
	tableName := table.Name.Source()
	if table.Db.Source() != "" {
		tableName = table.Db.Source() + "." + table.Name.Source()
	}
	output, err := util.With("types").
		Parse(text).
		Execute(map[string]any{
			"upperStartCamelObject": table.Name.ToCamel(),
			"lowerStartCamelObject": stringx.From(table.Name.ToCamel()).Untitle(),
			"fields":                fieldsString,
			"tableName":             tableName,
		})
	if err != nil {
		return "", err
	}

	return output.String(), nil
}
