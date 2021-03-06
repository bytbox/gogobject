package main

import (
	"bytes"
	"io"
	"fmt"
	"strings"
	"text/template"
)

var printf func(string, ...interface{})

func PrinterTo(w io.Writer) func(string, ...interface{}) {
	return func(format string, args ...interface{}) {
		fmt.Fprintf(w, format, args...)
	}
}

func LowerCaseToCamelCase(name string) string {
	var out bytes.Buffer
	for _, word := range strings.Split(name, "_") {
		word = strings.ToLower(word)
		if subst, ok := GConfig.WordSubst[word]; ok {
			out.WriteString(subst)
			continue
		}

		if word == "" {
			out.WriteString("_")
			continue
		}
		out.WriteString(strings.ToUpper(word[0:1]))
		out.WriteString(word[1:])
	}
	return out.String()
}

func MustTemplate(tpl string) *template.Template {
	tpl = strings.TrimSpace(tpl)
	return template.Must(
		template.New("").
			Delims("[<", ">]").
			Parse(tpl),
	)
}

func ExecuteTemplate(tpl *template.Template, args interface{}) string {
	var out bytes.Buffer
	tpl.Execute(&out, args)
	return out.String()
}

func CtorSuffix(name string) string {
	if len(name) > 4 {
		return LowerCaseToCamelCase(name[4:])
	}
	return ""
}

func PrintLinesWithIndent(str string) string {
	var out bytes.Buffer
	if str == "" {
		return ""
	}
	for _, line := range strings.Split(str, "\n") {
		fmt.Fprintf(&out, "\t%s\n", line)
	}

	return out.String()
}

func MapListToMapMap(maplist map[string][]string) map[string]map[string]bool {
	out := make(map[string]map[string]bool)
	for section, list := range maplist {
		out[section] = ListToMap(list)
	}
	return out
}

func ListToMap(list []string) map[string]bool {
	m := make(map[string]bool)
	for _, entry := range list {
		m[entry] = true
	}
	return m
}
