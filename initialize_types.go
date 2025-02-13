package mcp_golang

import (
	"encoding/json"
	"fmt"
)

// Capabilities a client may support. Known capabilities are defined here, in this
// schema, but this is not a closed set: any client can define its own, additional
// capabilities.
type ClientCapabilities struct {
	// Experimental, non-standard capabilities that the client supports.
	Experimental ClientCapabilitiesExperimental `json:"experimental,omitempty" yaml:"experimental,omitempty" mapstructure:"experimental,omitempty"`

	// Present if the client supports listing roots.
	Roots *ClientCapabilitiesRoots `json:"roots,omitempty" yaml:"roots,omitempty" mapstructure:"roots,omitempty"`

	// Present if the client supports sampling from an LLM.
	Sampling ClientCapabilitiesSampling `json:"sampling,omitempty" yaml:"sampling,omitempty" mapstructure:"sampling,omitempty"`
}

// Experimental, non-standard capabilities that the client supports.
type ClientCapabilitiesExperimental map[string]map[string]interface{}

// Present if the client supports listing roots.
type ClientCapabilitiesRoots struct {
	// Whether the client supports notifications for changes to the roots list.
	ListChanged *bool `json:"listChanged,omitempty" yaml:"listChanged,omitempty" mapstructure:"listChanged,omitempty"`
}

// Present if the client supports sampling from an LLM.
type ClientCapabilitiesSampling map[string]interface{}

type initializeRequestParams struct {
	// Capabilities corresponds to the JSON schema field "capabilities".
	Capabilities ClientCapabilities `json:"capabilities" yaml:"capabilities" mapstructure:"capabilities"`

	// ClientInfo corresponds to the JSON schema field "clientInfo".
	ClientInfo implementation `json:"clientInfo" yaml:"clientInfo" mapstructure:"clientInfo"`

	// The latest version of the Model Context Protocol that the client supports. The
	// client MAY decide to support older versions as well.
	ProtocolVersion string `json:"protocolVersion" yaml:"protocolVersion" mapstructure:"protocolVersion"`
}

// Capabilities that a server may support. Known capabilities are defined here, in
// this schema, but this is not a closed set: any server can define its own,
// additional capabilities.
type ServerCapabilities struct {
	// Experimental, non-standard capabilities that the server supports.
	Experimental ServerCapabilitiesExperimental `json:"experimental,omitempty" yaml:"experimental,omitempty" mapstructure:"experimental,omitempty"`

	// Present if the server supports sending log messages to the client.
	Logging ServerCapabilitiesLogging `json:"logging,omitempty" yaml:"logging,omitempty" mapstructure:"logging,omitempty"`

	// Present if the server offers any prompt templates.
	Prompts *ServerCapabilitiesPrompts `json:"prompts,omitempty" yaml:"prompts,omitempty" mapstructure:"prompts,omitempty"`

	// Present if the server offers any resources to read.
	Resources *ServerCapabilitiesResources `json:"resources,omitempty" yaml:"resources,omitempty" mapstructure:"resources,omitempty"`

	// Present if the server offers any tools to call.
	Tools *ServerCapabilitiesTools `json:"tools,omitempty" yaml:"tools,omitempty" mapstructure:"tools,omitempty"`
}

// Experimental, non-standard capabilities that the server supports.
type ServerCapabilitiesExperimental map[string]map[string]interface{}

// Present if the server supports sending log messages to the client.
type ServerCapabilitiesLogging map[string]interface{}

// Present if the server offers any prompt templates.
type ServerCapabilitiesPrompts struct {
	// Whether this server supports notifications for changes to the prompt list.
	ListChanged *bool `json:"listChanged,omitempty" yaml:"listChanged,omitempty" mapstructure:"listChanged,omitempty"`
}

// Present if the server offers any resources to read.
type ServerCapabilitiesResources struct {
	// Whether this server supports notifications for changes to the resource list.
	ListChanged *bool `json:"listChanged,omitempty" yaml:"listChanged,omitempty" mapstructure:"listChanged,omitempty"`

	// Whether this server supports subscribing to resource updates.
	Subscribe *bool `json:"subscribe,omitempty" yaml:"subscribe,omitempty" mapstructure:"subscribe,omitempty"`
}

// Present if the server offers any tools to call.
type ServerCapabilitiesTools struct {
	// Whether this server supports notifications for changes to the tool list.
	ListChanged *bool `json:"listChanged,omitempty" yaml:"listChanged,omitempty" mapstructure:"listChanged,omitempty"`
}

// After receiving an initialize request from the client, the server sends this
// response.
type InitializeResponse struct {
	// This result property is reserved by the protocol to allow clients and servers
	// to attach additional metadata to their responses.
	Meta initializeResultMeta `json:"_meta,omitempty" yaml:"_meta,omitempty" mapstructure:"_meta,omitempty"`

	// Capabilities corresponds to the JSON schema field "capabilities".
	Capabilities ServerCapabilities `json:"capabilities" yaml:"capabilities" mapstructure:"capabilities"`

	// Instructions describing how to use the server and its features.
	//
	// This can be used by clients to improve the LLM's understanding of available
	// tools, resources, etc. It can be thought of like a "hint" to the model. For
	// example, this information MAY be added to the system prompt.
	Instructions *string `json:"instructions,omitempty" yaml:"instructions,omitempty" mapstructure:"instructions,omitempty"`

	// The version of the Model Context Protocol that the server wants to use. This
	// may not match the version that the client requested. If the client cannot
	// support this version, it MUST disconnect.
	ProtocolVersion string `json:"protocolVersion" yaml:"protocolVersion" mapstructure:"protocolVersion"`

	// ServerInfo corresponds to the JSON schema field "serverInfo".
	ServerInfo implementation `json:"serverInfo" yaml:"serverInfo" mapstructure:"serverInfo"`
}

// This result property is reserved by the protocol to allow clients and servers to
// attach additional metadata to their responses.
type initializeResultMeta map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *InitializeResponse) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["capabilities"]; raw != nil && !ok {
		return fmt.Errorf("field capabilities in initializeResult: required")
	}
	if _, ok := raw["protocolVersion"]; raw != nil && !ok {
		return fmt.Errorf("field protocolVersion in initializeResult: required")
	}
	if _, ok := raw["serverInfo"]; raw != nil && !ok {
		return fmt.Errorf("field serverInfo in initializeResult: required")
	}
	type Plain InitializeResponse
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = InitializeResponse(plain)
	return nil
}

// Describes the name and version of an MCP implementation.
type implementation struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// Version corresponds to the JSON schema field "version".
	Version string `json:"version" yaml:"version" mapstructure:"version"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *implementation) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["name"]; raw != nil && !ok {
		return fmt.Errorf("field name in implementation: required")
	}
	if _, ok := raw["version"]; raw != nil && !ok {
		return fmt.Errorf("field version in implementation: required")
	}
	type Plain implementation
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = implementation(plain)
	return nil
}
