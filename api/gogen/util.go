package gogen

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/cat3306/ginctl/api/spec"
	"github.com/cat3306/ginctl/api/util"
	"github.com/cat3306/ginctl/pkg/golang"
	"github.com/cat3306/ginctl/util/pathx"
	"github.com/zeromicro/go-zero/core/collection"
)

type fileGenConfig struct {
	dir             string
	subdir          string
	filename        string
	templateName    string
	category        string
	templateFile    string
	builtinTemplate string
	data            any
}

func genFile(c fileGenConfig) error {
	fp, created, err := util.MaybeCreateFile(c.dir, c.subdir, c.filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	var text string
	if len(c.category) == 0 || len(c.templateFile) == 0 {
		text = c.builtinTemplate
	} else {
		text, err = pathx.LoadTemplate(c.category, c.templateFile, c.builtinTemplate)
		if err != nil {
			return err
		}
	}

	t := template.Must(template.New(c.templateName).Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.data)
	if err != nil {
		return err
	}

	code := golang.FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func writeProperty(writer io.Writer, name, tag, comment string, tp spec.Type, indent int) error {
	util.WriteIndent(writer, indent)
	var err error
	if len(comment) > 0 {
		comment = strings.TrimPrefix(comment, "//")
		comment = "//" + comment
		_, err = fmt.Fprintf(writer, "%s %s %s %s\n", strings.Title(name), tp.Name(), tag, comment)
	} else {
		_, err = fmt.Fprintf(writer, "%s %s %s\n", strings.Title(name), tp.Name(), tag)
	}

	return err
}

func getMiddleware(api *spec.ApiSpec) []string {
	result := collection.NewSet()
	uniq := make(map[string]struct{})
	for _, g := range api.Service.Groups {
		middleware := g.GetAnnotation("middleware")
		if len(middleware) > 0 {
			for _, item := range strings.Split(middleware, ",") {
				item = strings.TrimSpace(item)
				_, ok := uniq[item]
				if !ok {
					result.Add(item)
					uniq[item] = struct{}{}
				}
			}
		}
	}

	return result.KeysStr()
}

func responseGoTypeName(r spec.Route, pkg ...string) string {
	if r.ResponseType == nil {
		return ""
	}

	resp := golangExpr(r.ResponseType, pkg...)
	switch r.ResponseType.(type) {
	case spec.DefineStruct:
		if !strings.HasPrefix(resp, "*") {
			return "*" + resp
		}
	}

	return resp
}

func requestGoTypeName(r spec.Route, pkg ...string) string {
	if r.RequestType == nil {
		return ""
	}

	return golangExpr(r.RequestType, pkg...)
}

func golangExpr(ty spec.Type, pkg ...string) string {
	switch v := ty.(type) {
	case spec.PrimitiveType:
		return v.RawName
	case spec.DefineStruct:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("%s.%s", pkg[0], strings.Title(v.RawName))
	case spec.ArrayType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("[]%s", golangExpr(v.Value, pkg...))
	case spec.MapType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("map[%s]%s", v.Key, golangExpr(v.Value, pkg...))
	case spec.PointerType:
		if len(pkg) > 1 {
			panic("package cannot be more than 1")
		}

		if len(pkg) == 0 {
			return v.RawName
		}

		return fmt.Sprintf("*%s", golangExpr(v.Type, pkg...))
	case spec.InterfaceType:
		return v.RawName
	}

	return ""
}

func StrFirstLetterUp(s string) string {
	if len(s) <= 1 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
