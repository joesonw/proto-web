package plugin

import (
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	openapi_pb "github.com/joesonw/proto-web/pbgo/openapi"
)

func (p *Plugin) genFile(file *protogen.File) (map[string]H, H, error) {
	pathsH := map[string]H{}
	channelsH := H{}
	for _, service := range file.Services {
		paths, channels, err := p.genService(service)
		if err != nil {
			return nil, nil, err
		}
		for k, v := range paths {
			pathsH[k] = v
		}
		for k, v := range channels {
			channelsH[k] = v
		}
	}
	return pathsH, channelsH, nil
}

func (p *Plugin) genService(service *protogen.Service) (map[string]H, H, error) {
	prefix := ""
	serviceOptions := service.Desc.Options()
	if proto.HasExtension(serviceOptions, openapi_pb.E_Prefix) {
		prefix = proto.GetExtension(serviceOptions, openapi_pb.E_Prefix).(string)
	}
	paths := map[string]H{}
	channels := H{}
	for _, method := range service.Methods {
		httpMethod, path, err := getMethodAndPath(method)
		if err != nil {
			return nil, nil, err
		}
		if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
			h, err := p.genWebSocketMethod(method)
			if err != nil {
				return nil, nil, err
			}
			channels[prefix+path] = h
		} else {
			if httpMethod == http.MethodConnect {
				return nil, nil, fmt.Errorf("no stream allowed for unary method")
			}
			h, err := p.genHttpMethod(method)
			if err != nil {
				return nil, nil, err
			}
			paths[prefix+path] = H{
				strings.ToLower(httpMethod): h,
			}
		}
	}
	return paths, channels, nil
}

func (p *Plugin) genHttpMethod(method *protogen.Method) (H, error) {
	path := proto.GetExtension(method.Desc.Options(), openapi_pb.E_Path).(*openapi_pb.Path)
	httpMethod, _, err := getMethodAndPath(method)
	if err != nil {
		return nil, err
	}
	h := H{}
	h["description"] = commentSetToString(method.Comments)
	h["summary"] = path.GetSummary()
	h["operationId"] = path.GetId()
	parameters, requestBody := p.messageToRequest(httpMethod == http.MethodPost || httpMethod == http.MethodPut || httpMethod == http.MethodPatch, method.Input)
	h["parameters"] = parameters
	h["requestBody"] = requestBody
	h["responses"] = H{
		"default": p.messageToResponse(method.Output),
	}

	return h, nil
}

func (p *Plugin) genWebSocketMethod(method *protogen.Method) (H, error) {
	h := H{}
	h["publish"] = H{
		"message": H{
			"payload": p.messageToSchema(method.Input),
		},
	}
	h["subscribe"] = H{
		"message": H{
			"payload": p.messageToSchema(method.Output),
		},
	}
	return h, nil
}

func getMethodAndPath(method *protogen.Method) (string, string, error) {
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
