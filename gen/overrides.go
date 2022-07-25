package main

func NewOverrides() Overrides {
	var overrides Overrides = make(map[string]map[string]map[string]Field)

	// Possible string->float overrides. Uncomment which are needed.
	//overrides.Add("ce", "component", "analysisId", &FloatField{name: "analysisId"})
	//overrides.Add("ce", "task", "analysisId", &FloatField{name: "analysisId"})
	//overrides.Add("user_tokens", "generate", "token", &FloatField{"token"})
	//overrides.Add("components", "show", "version", &FloatField{name: "version"})
	//overrides.Add("measures", "component", "value", &FloatField{name: "value"})
	//overrides.Add("measures", "component_tree", "value", &FloatField{name: "value"})
	//overrides.Add("measures", "search_history", "value", &FloatField{name: "value"})
	//overrides.Add("metrics", "search", "id", &FloatField{name: "id"})
	//overrides.Add("project_links", "search", "id", &FloatField{name: "id"})
	//overrides.Add("project_pull_requests", "list", "key", &FloatField{name: "key"})
	//overrides.Add("qualitygates", "create", "error", &FloatField{name: "error"})
	//overrides.Add("qualitygates", "create", "warning", &FloatField{name: "warning"})
	overrides.Add("qualitygates", "get_by_project", "id", &FloatField{name: "id"})
	//overrides.Add("qualitygates", "list", "error", &FloatField{name: "error"})
	//overrides.Add("qualitygates", "project_status", "actualValue", &FloatField{name: "actualValue"})
	//overrides.Add("qualitygates", "project_status", "errorThreshold", &FloatField{name: "errorThreshold"})
	//overrides.Add("qualitygates", "show", "error", &FloatField{name: "error"})

	// TODO: the actives field of the search response should not be a struct field, but a map[string]SomeStruct field, which is not supported yet.
	//overrides.Add("rules", "search", "value", &FloatField{name: "value"})
	//overrides.Add("rules", "search", "defaultValue", &FloatField{name: "defaultValue"})

	//overrides.Add("rules", "show", "value", &FloatField{name: "value"})
	//overrides.Add("rules", "show", "defaultValue", &FloatField{name: "defaultValue"})
	//overrides.Add("settings", "show", "defaultValue", &FloatField{name: "defaultValue"})
	overrides.Add("user_groups", "create", "id", &FloatField{name: "id"})
	overrides.Add("user_groups", "search", "id", &FloatField{name: "id"})
	//overrides.Add("user_tokens", "generate", "token", &FloatField{name: "token"})

	return overrides
}

type Overrides map[string]map[string]map[string]Field

func (o Overrides) Add(endpoint string, actionKey string, name string, override Field) {
	if _, ok := o[endpoint]; !ok {
		o[endpoint] = make(map[string]map[string]Field)
	}
	if _, ok := o[endpoint][actionKey]; !ok {
		o[endpoint][actionKey] = map[string]Field{}
	}
	o[endpoint][actionKey][name] = override

}

func (o Overrides) Filter(endpoint string, actionKey string) map[string]Field {
	if endpointEntries, ok := o[endpoint]; ok {
		if actionEntries, ok := endpointEntries[actionKey]; ok {
			return actionEntries
		}
	}
	return map[string]Field{}
}
