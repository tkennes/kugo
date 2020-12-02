package kugo_model

type Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaProps struct {
	_ref                                     string                                                                            `json:"$ref,omitempty"`
	_schema                                  string                                                                            `json:"$schema,omitempty"`
	AdditionalItems                          *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaPropsOrBool  `json:"additionalItems,omitempty"`
	AdditionalProperties                     *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaPropsOrBool  `json:"additionalProperties,omitempty"`
	AllOf                                    []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaProps       `json:"allOf,omitempty"`
	AnyOf                                    []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaProps       `json:"anyOf,omitempty"`
	Default                                  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSON                   `json:"default,omitempty"`
	Definitions                              *interface{}                                                                      `json:"definitions,omitempty"`
	Dependencies                             *interface{}                                                                      `json:"dependencies,omitempty"`
	Description                              string                                                                            `json:"description,omitempty"`
	Enum                                     []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSON                  `json:"enum,omitempty"`
	Example                                  *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSON                   `json:"example,omitempty"`
	ExclusiveMaximum                         bool                                                                              `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum                         bool                                                                              `json:"exclusiveMinimum,omitempty"`
	ExternalDocs                             *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_ExternalDocumentation  `json:"externalDocs,omitempty"`
	Format                                   string                                                                            `json:"format,omitempty"`
	Id                                       string                                                                            `json:"id,omitempty"`
	Items                                    *Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaPropsOrArray `json:"items,omitempty"`
	MaxItems                                 int                                                                               `json:"maxItems,omitempty"`
	MaxLength                                int                                                                               `json:"maxLength,omitempty"`
	MaxProperties                            int                                                                               `json:"maxProperties,omitempty"`
	Maximum                                  float64                                                                           `json:"maximum,omitempty"`
	MinItems                                 int                                                                               `json:"minItems,omitempty"`
	MinLength                                int                                                                               `json:"minLength,omitempty"`
	MinProperties                            int                                                                               `json:"minProperties,omitempty"`
	Minimum                                  float64                                                                           `json:"minimum,omitempty"`
	MultipleOf                               float64                                                                           `json:"multipleOf,omitempty"`
	Not                                      **Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaProps       `json:"not,omitempty"`
	Nullable                                 bool                                                                              `json:"nullable,omitempty"`
	OneOf                                    []Io_k8s_apiextensions__apiserver_pkg_apis_apiextensions_v1_JSONSchemaProps       `json:"oneOf,omitempty"`
	Pattern                                  string                                                                            `json:"pattern,omitempty"`
	PatternProperties                        *interface{}                                                                      `json:"patternProperties,omitempty"`
	Properties                               *interface{}                                                                      `json:"properties,omitempty"`
	Required                                 []string                                                                          `json:"required,omitempty"`
	Title                                    string                                                                            `json:"title,omitempty"`
	Type                                     string                                                                            `json:"type,omitempty"`
	UniqueItems                              bool                                                                              `json:"uniqueItems,omitempty"`
	X__kubernetes__embedded__resource        bool                                                                              `json:"x-kubernetes-embedded-resource,omitempty"`
	X__kubernetes__int__or__string           bool                                                                              `json:"x-kubernetes-int-or-string,omitempty"`
	X__kubernetes__list__map__keys           []string                                                                          `json:"x-kubernetes-list-map-keys,omitempty"`
	X__kubernetes__list__type                string                                                                            `json:"x-kubernetes-list-type,omitempty"`
	X__kubernetes__map__type                 string                                                                            `json:"x-kubernetes-map-type,omitempty"`
	X__kubernetes__preserve__unknown__fields bool                                                                              `json:"x-kubernetes-preserve-unknown-fields,omitempty"`
}

