package plugin

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	openapi_pb "github.com/joesonw/proto-web/pbgo/openapi"
)

func (p *Plugin) messageToRequest(hasBody bool, message *protogen.Message) (parameters []H, requestBody H) {
	schema := H{}
	for _, field := range message.Fields {
		options := field.Desc.Options()
		name := ""
		in := ""
		if proto.HasExtension(options, openapi_pb.E_InHeader) {
			name = proto.GetExtension(options, openapi_pb.E_InHeader).(string)
			in = "header"
		} else if proto.HasExtension(options, openapi_pb.E_InPath) {
			name = proto.GetExtension(options, openapi_pb.E_InPath).(string)
			in = "path"
		} else if proto.HasExtension(options, openapi_pb.E_InCookie) {
			name = proto.GetExtension(options, openapi_pb.E_InCookie).(string)
			in = "cookie"
		} else if proto.HasExtension(options, openapi_pb.E_InQuery) || !hasBody {
			name = proto.GetExtension(options, openapi_pb.E_InQuery).(string)
			in = "query"
		} else {
			schema[string(field.Desc.Name())] = p.fieldToSchema(field)
		}
		if in == "" {
			continue
		}
		parameters = append(parameters, H{
			"name":            name,
			"in":              in,
			"required":        proto.GetExtension(options, openapi_pb.E_Required).(bool),
			"deprecated":      proto.GetExtension(options, openapi_pb.E_Deprecated).(bool),
			"allowEmptyValue": proto.GetExtension(options, openapi_pb.E_AllowEmptyValue).(bool),
			"description":     commentSetToString(field.Comments),
			"schema":          p.fieldToSchema(field),
		})
	}

	if hasBody {
		requestBody = H{
			"content": H{
				"application/json": H{
					"schema": H{
						"type":       "object",
						"properties": schema,
					},
				},
			},
			"description": commentSetToString(message.Comments),
		}
	}
	return
}

func (p *Plugin) messageToResponse(message *protogen.Message) H {
	schema := H{}
	headers := H{}
	for _, field := range message.Fields {
		options := field.Desc.Options()
		if proto.HasExtension(options, openapi_pb.E_InHeader) {
			headers[proto.GetExtension(options, openapi_pb.E_InHeader).(string)] = H{
				"description": commentSetToString(field.Comments),
				"schema":      p.fieldToSchema(field),
			}
		} else if proto.HasExtension(options, openapi_pb.E_InPath) {
		} else if proto.HasExtension(options, openapi_pb.E_InCookie) {
		} else if proto.HasExtension(options, openapi_pb.E_InQuery) {
		} else {
			schema[string(field.Desc.Name())] = p.fieldToSchema(field)
		}
	}

	return H{
		"description": commentSetToString(message.Comments),
		"content": H{
			"application/json": H{
				"schema": H{
					"type":       "object",
					"properties": schema,
				},
			},
		},
		"headers": headers,
	}
}

func (p *Plugin) messageToSchema(message *protogen.Message) H {
	properties := H{}
	for _, field := range message.Fields {
		properties[string(field.Desc.Name())] = p.fieldToSchema(field)
	}
	return H{
		"type":        "object",
		"description": commentSetToString(message.Comments),
		"properties":  properties,
	}
}

func (p *Plugin) fieldToSchema(field *protogen.Field) H {
	output := H{}
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		output = H{
			"type":        "boolean",
			"description": commentSetToString(field.Comments),
		}
	case protoreflect.EnumKind:
		if p.EnumAsString {
			enum := make([]string, len(field.Enum.Values))
			for i, e := range field.Enum.Values {
				enum[i] = string(e.Desc.Name())
			}
			output = H{
				"type":        "string",
				"description": commentSetToString(field.Comments),
				"enum":        enum,
			}
		} else {
			enum := make([]int32, len(field.Enum.Values))
			for i, e := range field.Enum.Values {
				enum[i] = int32(e.Desc.Number())
			}
			output = H{
				"type":        "integer",
				"description": commentSetToString(field.Comments),
				"enum":        enum,
			}
		}
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind:
		output = H{
			"type":        "integer",
			"description": commentSetToString(field.Comments),
		}
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		output = H{
			"type":        "number",
			"description": commentSetToString(field.Comments),
		}
	case protoreflect.StringKind:
		output = H{
			"type":        "string",
			"description": commentSetToString(field.Comments),
		}
	case protoreflect.BytesKind:
		panic(fmt.Sprintf("field type %s is not currently supported, please use %s", protoreflect.BytesKind.String(), protoreflect.StringKind.String()))
	case protoreflect.MessageKind:
		output = p.messageToSchema(field.Message)
	default:
		panic(fmt.Sprintf("field type %s is not currently supported", field.Desc.Kind().String()))
	}

	if field.Desc.IsList() {
		output = H{
			"type":        "array",
			"description": commentSetToString(field.Comments),
			"items":       output,
		}
	} else if field.Desc.IsMap() {
		output = H{
			"type":                 "object",
			"description":          commentSetToString(field.Comments),
			"additionalProperties": output,
		}
	}
	return output
}
