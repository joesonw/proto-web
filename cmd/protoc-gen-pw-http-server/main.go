package main

import (
	"flag"

	"github.com/joesonw/proto-tools/pkg/genutil"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/joesonw/proto-web/cmd/protoc-gen-pw-http-server/plugin"
	protoutil2 "github.com/joesonw/proto-web/pkg/protoutil"
)

func main() {
	var (
		flags        flag.FlagSet
		importPrefix = flags.String("import_prefix", "", "prefix to prepend to import paths")
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
		plg := &plugin.Plugin{}
		for _, file := range gen.Files {
			if !file.Generate {
				continue
			}
			filename := file.GeneratedFilenamePrefix + "_http_server.pb.go"
			g := genutil.New(gen.NewGeneratedFile(filename, file.GoImportPath), func(g *genutil.G) genutil.Generator {
				return protoutil2.GeneratorFunc(func() error {
					return plg.GenFile(file, g)
				})
			})
			if err := g.Generate(); err != nil {
				return err
			}
		}
		return nil
	})
}
