package plugin

import (
	"github.com/joesonw/proto-tools/pkg/genutil"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *Plugin) GenUnary(method *protogen.Method, options ServiceOptions, g *genutil.G) error {
	g.F("func _%s_%s_HttpHandler(srv interface{}, w %s, r *%s, params %s, interceptor %s) (interface{}, error) {", method.Parent.GoName, method.GoName, pkgHttp.Ident("ResponseWriter"), pkgHttp.Ident("Request"), pkgHttpRouter.Ident("Params"), pkgGrpc.Ident("UnaryServerInterceptor"))
	g.P("var err error")
	g.P("ctx := r.Context()")
	g.F("req := &%s{}", method.Input.GoIdent)
	if err := p.genUnaryRequestHandle(method, g); err != nil {
		return err
	}

	g.F("var res *%s", method.Output.GoIdent)
	g.P("if interceptor != nil {")
	g.F("res, err = srv.(%sServer).%s(ctx, req)", method.Parent.GoName, method.GoName)
	g.P("} else {")
	g.F("info := &%s{", pkgGrpc.Ident("UnaryServerInfo"))
	g.P("Server: srv,")
	g.F("FullMethod: \"%s\",", method.Desc.FullName())
	g.P("}")
	g.P("")

	g.F("handler := func (ctx %s, in interface{}) (interface{}, error) {", pkgContext.Ident("Context"))
	g.F("return srv.(%sServer).%s(ctx, in.(*%s))", method.Parent.GoName, method.GoName, method.Input.GoIdent)
	g.F("}")
	g.P("var resp interface{}")
	g.P("resp, err = interceptor(ctx, req, info, handler)")
	g.F("res = resp.(*%s)", method.Output.GoIdent)
	g.P("}")

	if err := p.genUnaryResponseHandle(method, g); err != nil {
		return err
	}
	g.P("}")
	return nil
}
