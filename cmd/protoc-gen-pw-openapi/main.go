package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/joesonw/proto-web/cmd/protoc-gen-pw-openapi/plugin"
)

func main() {
	var (
		flags        flag.FlagSet
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
		enumAsString = flags.Bool("enum_as_string", true, "if enum are processed as string")
	)
	importRewriteFunc := func(importPath protogen.GoImportPath) protogen.GoImportPath {
		switch importPath {
		case "context", "fmt", "math":
			return importPath
		}
		if *importPrefix != "" {
			return protogen.GoImportPath(*importPrefix) + importPath
		}
		return importPath
	}
	protogen.Options{
		ParamFunc:         flags.Set,
		ImportRewriteFunc: importRewriteFunc,
	}.Run(func(gen *protogen.Plugin) error {
		plg := &plugin.Plugin{
			EnumAsString: *enumAsString,
		}
		var files []*protogen.File
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			files = append(files, file)
		}
		return plg.Generate(files, gen.NewGeneratedFile("openapi.json", protogen.GoImportPath("")))
	})
}
