package plugin

import (
	"encoding/json"
	"regexp"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	openapi_pb "github.com/joesonw/proto-web/pbgo/openapi"
)

var (
	pathParamRegex = regexp.MustCompile(`{[a-zA-Z][a-zA-Z0-9_]*}`)
)

type H map[string]interface{}

type Plugin struct {
	EnumAsString bool
}

func (p *Plugin) Generate(files []*protogen.File, g *protogen.GeneratedFile) error {
	var info *openapi_pb.Info
	var servers []*openapi_pb.Server
	var security []*openapi_pb.SecurityRequirement
	var tags []*openapi_pb.Tag
	var externalDocs *openapi_pb.ExternalDocumentation

	for _, file := range files {
		options := file.Desc.Options()
		if proto.HasExtension(options, openapi_pb.E_Info) {
			info = proto.GetExtension(options, openapi_pb.E_Info).(*openapi_pb.Info)
		}
		if proto.HasExtension(options, openapi_pb.E_Servers) {
			servers = proto.GetExtension(options, openapi_pb.E_Servers).([]*openapi_pb.Server)
		}
		if proto.HasExtension(options, openapi_pb.E_Security) {
			security = proto.GetExtension(options, openapi_pb.E_Security).([]*openapi_pb.SecurityRequirement)
		}
		if proto.HasExtension(options, openapi_pb.E_Tags) {
			tags = proto.GetExtension(options, openapi_pb.E_Tags).([]*openapi_pb.Tag)
		}
		if proto.HasExtension(options, openapi_pb.E_ExternalDocs) {
			externalDocs = proto.GetExtension(options, openapi_pb.E_ExternalDocs).(*openapi_pb.ExternalDocumentation)
		}
	}

	openapi := H{
		"openapi": "3.1",
	}
	if info != nil {
		infoH := H{
			"title":          info.GetTitle(),
			"summary":        info.GetSummary(),
			"description":    info.GetDescription(),
			"termsOfService": info.GetTermsOfService(),
			"version":        info.GetVersion(),
		}

		if contact := info.GetContact(); contact != nil {
			infoH["contact"] = H{
				"name":  contact.GetName(),
				"url":   contact.GetUrl(),
				"email": contact.GetEmail(),
			}
		}

		if license := info.GetLicense(); license != nil {
			infoH["license"] = H{
				"name":       license.GetName(),
				"identifier": license.GetIdentifier(),
				"url":        license.GetUrl(),
			}
		}

		openapi["info"] = infoH
	}

	var serversH []interface{}
	for _, server := range servers {
		serversH = append(serversH, serverToJson(server))
	}
	openapi["servers"] = serversH

	var securityH []interface{}
	for _, s := range security {
		h := H{}
		for k, v := range s.GetScopes() {
			h[k] = v.GetValues()
		}
		securityH = append(securityH, h)
	}
	openapi["security"] = securityH

	var tagsH []interface{}
	for _, tag := range tags {
		h := H{
			"name":         tag.GetName(),
			"description":  tag.GetDescription(),
			"externalDocs": externalDocsToJson(tag.GetExternalDocs()),
		}
		tagsH = append(tagsH, h)
	}
	openapi["tags"] = tagsH
	openapi["externalDocs"] = externalDocsToJson(externalDocs)

	pathsH := map[string]H{}
	channelsH := H{}
	for _, file := range files {
		paths, channels, err := p.genFile(file)
		if err != nil {
			return err
		}
		for k, methods := range paths {
			if _, ok := pathsH[k]; !ok {
				pathsH[k] = H{}
			}
			for method, v := range methods {
				pathsH[k][method] = v
			}
		}
		for k, v := range channels {
			channelsH[k] = v
		}
	}
	openapi["paths"] = pathsH
	openapi["channels"] = channelsH

	b, _ := json.MarshalIndent(openapi, "", strings.Repeat(" ", 4))
	if _, err := g.Write(b); err != nil {
		return err
	}

	return nil
}

func serverToJson(server *openapi_pb.Server) interface{} {
	variablesH := H{}
	for k, v := range server.GetVariables() {
		variablesH[k] = H{
			"enum":        v.GetEnum(),
			"default":     v.GetDefault(),
			"description": v.GetDescription(),
		}
	}
	h := H{
		"url":         server.GetUrl(),
		"description": server.GetDescription(),
		"variables":   variablesH,
	}
	return h
}

func externalDocsToJson(ed *openapi_pb.ExternalDocumentation) interface{} {
	if ed == nil {
		return nil
	}
	return H{
		"description": ed.GetDescription(),
		"url":         ed.GetUrl(),
	}
}

func commentSetToString(cs protogen.CommentSet) string {
	var s []string
	for _, c := range cs.LeadingDetached {
		s = append(s, c.String())
	}

	if cs.Leading != "" {
		s = append(s, cs.Leading.String())
	}

	if cs.Trailing != "" {
		s = append(s, cs.Trailing.String())
	}

	return strings.Join(s, "\n")
}
