package plugin

import (
	"fmt"
	"net/http"

	"github.com/joesonw/proto-tools/pkg/genutil"
	"github.com/joesonw/proto-tools/pkg/protoutil"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	openapi_pb "github.com/joesonw/proto-web/pbgo/openapi"
)

type ServiceOptions struct {
	prefix string
}

func (p *Plugin) GenService(service *protogen.Service, g *genutil.G) error {
	options := ServiceOptions{}
	protoOptions := service.Desc.Options()
	if proto.HasExtension(protoOptions, openapi_pb.E_Prefix) {
		options.prefix = proto.GetExtension(protoOptions, openapi_pb.E_Prefix).(string)
	}

	g.F("func Register%sHTTPServer(s %s, srv %sServer)  {", service.GoName, pkgProtoWeb.Ident("ServiceRegistrar"), service.GoName)
	g.F("s.RegisterService(&%s_HttpServiceDesc, srv)", service.GoName)
	g.P("}")

	for _, method := range service.Methods {
		if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
			break
		}
	}

	for _, method := range service.Methods {
		if !proto.HasExtension(method.Desc.Options(), openapi_pb.E_Path) {
			return fmt.Errorf("method %s of service %s does not have path annotation", method.Desc.Name(), service.Desc.Name())
		}

		isServer := method.Desc.IsStreamingServer()
		isClient := method.Desc.IsStreamingClient()

		if !isServer && !isClient {
			if err := p.GenUnary(method, options, g); err != nil {
				return err
			}
		} else {
			if err := p.GenStream(method, options, g); err != nil {
				return err
			}
		}
	}

	g.F("var %s_HttpServiceDesc = %s{", service.GoName, pkgProtoWeb.Ident("ServiceDesc"))
	g.F("ServiceName: \"%s\",", service.Desc.FullName())
	g.F("HandlerType: (*%sServer)(nil),", service.GoName)
	g.F("Methods: []%s{", pkgProtoWeb.Ident("MethodDesc"))
	for _, method := range service.Methods {
		isServer := method.Desc.IsStreamingServer()
		isClient := method.Desc.IsStreamingClient()
		if !isServer && !isClient {
			httpMethod, path, err := p.getMethodAndPath(method)
			if err != nil {
				return err
			}
			if httpMethod == http.MethodConnect {
				return fmt.Errorf("cannot have stream path for non-stream method")
			}
			g.P("{")
			g.F("MethodName: \"%s\",", method.Desc.Name())
			g.F("Path: \"%s\",", path)
			g.F("HttpMethod: \"%s\",", httpMethod)
			g.F("Handler: _%s_%s_HttpHandler,", service.GoName, method.GoName)
			g.P("},")
		}
	}
	g.P("},")
	g.F("Streams: []%s{", pkgProtoWeb.Ident("StreamDesc"))
	for _, method := range service.Methods {
		isServer := method.Desc.IsStreamingServer()
		isClient := method.Desc.IsStreamingClient()
		_, path, err := p.getMethodAndPath(method)
		if err != nil {
			return err
		}
		if isServer || isClient {
			g.P("{")
			g.F("StreamName: \"%s\",", method.Desc.Name())
			g.F("Path: \"%s\",", path)
			g.F("Handler: _%s_%s_Handler,", service.GoName, method.GoName)
			if isServer {
				g.P("ServerStreams: true,")
			}
			if isClient {
				g.P("ClientStreams: true,")
			}
			g.P("},")
		}
	}
	g.P("},")
	g.F("Metadata: \"%s\",", service.Location.SourceFile)
	g.P("}")

	return nil
}

func (p *Plugin) validateUnaryRequest(method *protogen.Method, message *protogen.Message) error {
	return nil
}

func (p *Plugin) validateUnaryResponse(method *protogen.Method, message *protogen.Message) error {
	for _, field := range message.Fields {
		options := field.Desc.Options()
		if proto.GetExtension(options, openapi_pb.E_InPath).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in path, because it's a response message", field.Desc.Name(), message.Desc.Name())
		}
		if proto.GetExtension(options, openapi_pb.E_InQuery).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in query, because it's a response message", field.Desc.Name(), message.Desc.Name())
		}
	}
	return nil
}

func (p *Plugin) validateStreamMessage(method *protogen.Method, message *protogen.Message) error {
	for _, field := range message.Fields {
		options := field.Desc.Options()
		if proto.GetExtension(options, openapi_pb.E_InPath).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in path, because it's a stream message", field.Desc.Name(), message.Desc.Name())
		}
		if proto.GetExtension(options, openapi_pb.E_InQuery).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in path, because it's a stream message", field.Desc.Name(), message.Desc.Name())
		}
		if proto.GetExtension(options, openapi_pb.E_InHeader).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in header, because it's a stream message", field.Desc.Name(), message.Desc.Name())
		}
		if proto.GetExtension(options, openapi_pb.E_InCookie).(string) != "" {
			return fmt.Errorf("field %s of message %s cannot be in cookie, because it's a stream message", field.Desc.Name(), message.Desc.Name())
		}
	}
	return nil
}

func (p *Plugin) regulatePath(path string) string {
	indexes := pathParamRegex.FindAllStringIndex(path, -1)
	result := path
	offset := 0
	for _, pair := range indexes {
		begin := pair[0] - offset
		end := pair[1] - offset
		result = result[:begin] + ":" + result[begin+1:end-1] + result[end:]
		offset += 1
	}
	return result
}

func (p *Plugin) genUnaryRequestHandle(method *protogen.Method, g *genutil.G) error {
	httpMethod, _, err := p.getMethodAndPath(method)
	if err != nil {
		return err
	}
	hasBody := httpMethod == http.MethodPost || httpMethod == http.MethodPut
	if hasBody {
		g.F("b, err := %s(r.Body)", pkgIoutil.Ident("ReadAll"))
		g.P("_ = r.Body.Close()")
		g.P("if err != nil {")
		g.P("return nil, err")
		g.P("}")
		g.F("if err := (%s{}).Unmarshal(b, req); err != nil {", pkgProtojson.Ident("UnmarshalOptions"))
		g.P("return nil, err")
		g.P("}")
	}
	for i, field := range method.Input.Fields {
		options := field.Desc.Options()
		source := ""
		if proto.HasExtension(options, openapi_pb.E_InQuery) {
			source = fmt.Sprintf("r.URL.Query().Get(\"%s\")", proto.GetExtension(options, openapi_pb.E_InQuery).(string))
		} else if proto.HasExtension(options, openapi_pb.E_InPath) {
			source = fmt.Sprintf("params.ByName(\"%s\")", proto.GetExtension(options, openapi_pb.E_InPath).(string))
		} else if proto.HasExtension(options, openapi_pb.E_InHeader) {
			source = fmt.Sprintf("r.Header.Get(\"%s\")", proto.GetExtension(options, openapi_pb.E_InHeader).(string))
		} else if proto.HasExtension(options, openapi_pb.E_InCookie) {
			name := proto.GetExtension(options, openapi_pb.E_InCookie).(string)
			g.F("c%d, err := r.Cookie(\"%s\")", i+1, name)
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
			source = fmt.Sprintf("c%d", i+1)
		} else {
			if !hasBody {
				source = proto.GetExtension(options, openapi_pb.E_InQuery).(string)
			} else {
				continue
			}
		}
		checkError, err := p.genConvertFromString(field, source, fmt.Sprintf("req.%s", field.GoName), i+1, g)
		if err != nil {
			return err
		}
		if checkError {
			g.P("if err != nil {")
			g.P("return nil, err")
			g.P("}")
		}
	}

	return nil
}

func (p *Plugin) genUnaryResponseHandle(method *protogen.Method, g *genutil.G) error {
	for _, field := range method.Output.Fields {
		options := field.Desc.Options()
		if proto.HasExtension(options, openapi_pb.E_InQuery) {
			return fmt.Errorf("field %s of message %s: cannot have in_query annotation in response", field.Desc.Name(), method.Input.Desc.Name())
		} else if proto.HasExtension(options, openapi_pb.E_InPath) {
			return fmt.Errorf("field %s of message %s: cannot have in_path annotation in response", field.Desc.Name(), method.Input.Desc.Name())
		} else if proto.HasExtension(options, openapi_pb.E_InHeader) {
			g.F("w.Header().Set(\"%s\", %s(\"%%v\", res.%s))", proto.GetExtension(options, openapi_pb.E_InHeader).(string), pkgFmt.Ident("Sprintf"), field.GoName)
			g.F("res.%s = %s", field.GoName, field.Desc.Default().String())
		} else if proto.HasExtension(options, openapi_pb.E_InCookie) {
			return fmt.Errorf("field %s of message %s: cannot have in_cookie annotation in response", field.Desc.Name(), method.Input.Desc.Name())
		}
	}
	g.F("return res, nil")
	return nil
}

func (p *Plugin) getMethodAndPath(method *protogen.Method) (string, string, error) {
	path := proto.GetExtension(method.Desc.Options(), openapi_pb.E_Path).(*openapi_pb.Path)
	if path.Get != "" {
		return http.MethodGet, path.Get, nil
	} else if path.Put != "" {
		return http.MethodPut, path.Put, nil
	} else if path.Post != "" {
		return http.MethodPost, path.Post, nil
	} else if path.Delete != "" {
		return http.MethodDelete, path.Delete, nil
	} else if path.Options != "" {
		return http.MethodOptions, path.Options, nil
	} else if path.Head != "" {
		return http.MethodHead, path.Head, nil
	} else if path.Patch != "" {
		return http.MethodPatch, path.Patch, nil
	} else if path.Trace != "" {
		return http.MethodTrace, path.Trace, nil
	} else if path.Stream != "" {
		return http.MethodConnect, path.Stream, nil
	} else {
		return "", "", fmt.Errorf("no valid method found for method %s of service %s", method.Desc.Name(), method.Parent.Desc.Name())
	}
}

func (p *Plugin) genConvertFromString(field *protogen.Field, source, target string, index int, g *genutil.G) (bool, error) {
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		g.F("%s = %s", target, source)
		return false, nil
	case protoreflect.BoolKind:
		g.F("%s, err = %s(%s)", target, pkgStrconv.Ident("ParseBool"), source)
		return true, nil
	case protoreflect.EnumKind:
		g.F("%s = %s(%s_value[%s])", target, field.Enum.GoIdent, field.Enum.GoIdent, source)
		return false, nil
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		g.F("x%d, err := %s(%s, 10, 64)", index, pkgStrconv.Ident("ParseInt"), source)
		g.F("%s = %s(x%d)", target, protoutil.FieldGoType(g.Q, field), index)
		return true, nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		g.F("x%d, err := %s(%s,  64)", index, pkgStrconv.Ident("ParseFloat"), source)
		g.F("%s = %s(x%d)", target, protoutil.FieldGoType(g.Q, field), index)
		return true, nil
	default:
		return false, fmt.Errorf("field %s(type %s) cannot be casted from string", field.Desc.Name(), field.Desc.Kind().String())
	}
}
