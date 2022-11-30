package rules

import paging "github.com/reinoudk/go-sonarcloud/sonarcloud/paging"

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// RepositoriesRequest List available rule repositories
type RepositoriesRequest struct {
	Language string `form:"language,omitempty"` // A language key; if provided, only repositories for the given language will be returned
	Q        string `form:"q,omitempty"`        // A pattern to match repository keys/names against
}

// RepositoriesResponse is the response for RepositoriesRequest
type RepositoriesResponse struct {
	Repositories []struct {
		Key      string `json:"key,omitempty"`
		Language string `json:"language,omitempty"`
		Name     string `json:"name,omitempty"`
	} `json:"repositories,omitempty"`
}

// SearchRequest Search for a collection of relevant rules matching a specified query.<br/>Since 5.5, following fields in the response have been deprecated :<ul><li>"effortToFixDescription" becomes "gapDescription"</li><li>"debtRemFnCoeff" becomes "remFnGapMultiplier"</li><li>"defaultDebtRemFnCoeff" becomes "defaultRemFnGapMultiplier"</li><li>"debtRemFnOffset" becomes "remFnBaseEffort"</li><li>"defaultDebtRemFnOffset" becomes "defaultRemFnBaseEffort"</li><li>"debtOverloaded" becomes "remFnOverloaded"</li></ul>
type SearchRequest struct {
	Activation          string `form:"activation,omitempty"`          // Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.
	ActiveSeverities    string `form:"active_severities,omitempty"`   // Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.
	Asc                 string `form:"asc,omitempty"`                 // Ascending sort
	AvailableSince      string `form:"available_since,omitempty"`     // Filters rules added since date. Format is yyyy-MM-dd
	Cwe                 string `form:"cwe,omitempty"`                 // Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.
	F                   string `form:"f,omitempty"`                   // Comma-separated list of the fields to be returned in response. All the fields are returned by default, except actives.Since 5.5, following fields have been deprecated :<ul><li>"defaultDebtRemFn" becomes "defaultRemFn"</li><li>"debtRemFn" becomes "remFn"</li><li>"effortToFixDescription" becomes "gapDescription"</li><li>"debtOverloaded" becomes "remFnOverloaded"</li></ul>
	Facets              string `form:"facets,omitempty"`              // Comma-separated list of the facets to be computed. No facet is computed by default.
	IncludeExternal     string `form:"include_external,omitempty"`    // Include external engine rules in the results
	Inheritance         string `form:"inheritance,omitempty"`         // Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.
	IsTemplate          string `form:"is_template,omitempty"`         // Filter template rules
	Languages           string `form:"languages,omitempty"`           // Comma-separated list of languages
	Organization        string `form:"organization,omitempty"`        // Organization key
	OwaspTop10          string `form:"owaspTop10,omitempty"`          // Comma-separated list of OWASP Top 10 lowercase categories.
	Q                   string `form:"q,omitempty"`                   // UTF-8 search query
	Qprofile            string `form:"qprofile,omitempty"`            // Quality profile key to filter on. Used only if the parameter 'activation' is set.
	Repositories        string `form:"repositories,omitempty"`        // Comma-separated list of repositories
	RuleKey             string `form:"rule_key,omitempty"`            // Key of rule to search for
	RuleKeys            string `form:"rule_keys,omitempty"`           // Rule keys
	S                   string `form:"s,omitempty"`                   // Sort field
	SansTop25           string `form:"sansTop25,omitempty"`           // Comma-separated list of SANS Top 25 categories.
	Severities          string `form:"severities,omitempty"`          // Comma-separated list of default severities. Not the same than severity of rules in Quality profiles.
	SonarsourceSecurity string `form:"sonarsourceSecurity,omitempty"` // Comma-separated list of SonarSource security categories. Use 'others' to select rules not associated with any category
	Statuses            string `form:"statuses,omitempty"`            // Comma-separated list of status codes
	Tags                string `form:"tags,omitempty"`                // Comma-separated list of tags. Returned rules match any of the tags (OR operator)
	TemplateKey         string `form:"template_key,omitempty"`        // Key of the template rule to filter on. Used to search for the custom rules based on this template.
	Types               string `form:"types,omitempty"`               // Comma-separated list of types. Returned rules match any of the tags (OR operator)
}

// SearchResponse is the response for SearchRequest
type SearchResponse struct {
	Actives struct {
		SquidClassCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:ClassCyclomaticComplexity,omitempty"`
		SquidMethodCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:MethodCyclomaticComplexity,omitempty"`
		SquidS1067 []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:S1067,omitempty"`
	} `json:"actives,omitempty"`
	Facets []struct {
		Name   string `json:"name,omitempty"`
		Values []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	P     float64 `json:"p,omitempty"`
	Ps    float64 `json:"ps,omitempty"`
	Rules []struct {
		CreatedAt   string `json:"createdAt,omitempty"`
		HtmlDesc    string `json:"htmlDesc,omitempty"`
		InternalKey string `json:"internalKey,omitempty"`
		IsExternal  bool   `json:"isExternal,omitempty"`
		IsTemplate  bool   `json:"isTemplate,omitempty"`
		Key         string `json:"key,omitempty"`
		Lang        string `json:"lang,omitempty"`
		LangName    string `json:"langName,omitempty"`
		Name        string `json:"name,omitempty"`
		Params      []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		Repo        string   `json:"repo,omitempty"`
		Scope       string   `json:"scope,omitempty"`
		Severity    string   `json:"severity,omitempty"`
		Status      string   `json:"status,omitempty"`
		SysTags     []string `json:"sysTags,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdatedAt   string   `json:"updatedAt,omitempty"`
		HtmlNote    string   `json:"htmlNote,omitempty"`
		MdNote      string   `json:"mdNote,omitempty"`
		NoteLogin   string   `json:"noteLogin,omitempty"`
		TemplateKey string   `json:"templateKey,omitempty"`
	} `json:"rules,omitempty"`
	Total float64 `json:"total,omitempty"`
}

// GetPaging extracts the paging from SearchResponse
func (r *SearchResponse) GetPaging() *paging.Paging {
	return &paging.Paging{

		PageIndex: int(r.P),
		PageSize:  int(r.Ps),
		Total:     int(r.Total),
	}
}

// SearchResponseAll is the collection for SearchRequest
type SearchResponseAll struct {
	Actives struct {
		SquidClassCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:ClassCyclomaticComplexity,omitempty"`
		SquidMethodCyclomaticComplexity []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:MethodCyclomaticComplexity,omitempty"`
		SquidS1067 []struct {
			Inherit string `json:"inherit,omitempty"`
			Params  []struct {
				Key   string `json:"key,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"params,omitempty"`
			QProfile string `json:"qProfile,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"squid:S1067,omitempty"`
	} `json:"actives,omitempty"`
	Facets []struct {
		Name   string `json:"name,omitempty"`
		Values []struct {
			Count float64 `json:"count,omitempty"`
			Val   string  `json:"val,omitempty"`
		} `json:"values,omitempty"`
	} `json:"facets,omitempty"`
	Rules []struct {
		CreatedAt   string `json:"createdAt,omitempty"`
		HtmlDesc    string `json:"htmlDesc,omitempty"`
		InternalKey string `json:"internalKey,omitempty"`
		IsExternal  bool   `json:"isExternal,omitempty"`
		IsTemplate  bool   `json:"isTemplate,omitempty"`
		Key         string `json:"key,omitempty"`
		Lang        string `json:"lang,omitempty"`
		LangName    string `json:"langName,omitempty"`
		Name        string `json:"name,omitempty"`
		Params      []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		Repo        string   `json:"repo,omitempty"`
		Scope       string   `json:"scope,omitempty"`
		Severity    string   `json:"severity,omitempty"`
		Status      string   `json:"status,omitempty"`
		SysTags     []string `json:"sysTags,omitempty"`
		Tags        []string `json:"tags,omitempty"`
		Type        string   `json:"type,omitempty"`
		UpdatedAt   string   `json:"updatedAt,omitempty"`
		HtmlNote    string   `json:"htmlNote,omitempty"`
		MdNote      string   `json:"mdNote,omitempty"`
		NoteLogin   string   `json:"noteLogin,omitempty"`
		TemplateKey string   `json:"templateKey,omitempty"`
	} `json:"rules,omitempty"`
}

// ShowRequest Get detailed information about a rule<br>Since 5.5, following fields in the response have been deprecated :<ul><li>"effortToFixDescription" becomes "gapDescription"</li><li>"debtRemFnCoeff" becomes "remFnGapMultiplier"</li><li>"defaultDebtRemFnCoeff" becomes "defaultRemFnGapMultiplier"</li><li>"debtRemFnOffset" becomes "remFnBaseEffort"</li><li>"defaultDebtRemFnOffset" becomes "defaultRemFnBaseEffort"</li><li>"debtOverloaded" becomes "remFnOverloaded"</li></ul>In 7.1, the field 'scope' has been added.
type ShowRequest struct {
	Actives      string `form:"actives,omitempty"`      // Show rule's activations for all profiles ("active rules")
	Key          string `form:"key,omitempty"`          // Rule key
	Organization string `form:"organization,omitempty"` // Organization key
}

// ShowResponse is the response for ShowRequest
type ShowResponse struct {
	Actives []struct {
		Inherit string `json:"inherit,omitempty"`
		Params  []struct {
			Key   string `json:"key,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"params,omitempty"`
		QProfile string `json:"qProfile,omitempty"`
		Severity string `json:"severity,omitempty"`
	} `json:"actives,omitempty"`
	Rule struct {
		DefaultRemFnBaseEffort    string `json:"defaultRemFnBaseEffort,omitempty"`
		DefaultRemFnGapMultiplier string `json:"defaultRemFnGapMultiplier,omitempty"`
		DefaultRemFnType          string `json:"defaultRemFnType,omitempty"`
		GapDescription            string `json:"gapDescription,omitempty"`
		HtmlDesc                  string `json:"htmlDesc,omitempty"`
		InternalKey               string `json:"internalKey,omitempty"`
		IsExternal                bool   `json:"isExternal,omitempty"`
		Key                       string `json:"key,omitempty"`
		Lang                      string `json:"lang,omitempty"`
		LangName                  string `json:"langName,omitempty"`
		Name                      string `json:"name,omitempty"`
		Params                    []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			Desc         string `json:"desc,omitempty"`
			Key          string `json:"key,omitempty"`
		} `json:"params,omitempty"`
		RemFnBaseEffort    string   `json:"remFnBaseEffort,omitempty"`
		RemFnGapMultiplier string   `json:"remFnGapMultiplier,omitempty"`
		RemFnOverloaded    bool     `json:"remFnOverloaded,omitempty"`
		RemFnType          string   `json:"remFnType,omitempty"`
		Repo               string   `json:"repo,omitempty"`
		Scope              string   `json:"scope,omitempty"`
		Severity           string   `json:"severity,omitempty"`
		Status             string   `json:"status,omitempty"`
		SysTags            []string `json:"sysTags,omitempty"`
		Tags               []string `json:"tags,omitempty"`
		Template           bool     `json:"template,omitempty"`
		Type               string   `json:"type,omitempty"`
	} `json:"rule,omitempty"`
}

// TagsRequest List rule tags
type TagsRequest struct {
	Organization string `form:"organization,omitempty"` // Organization key
	Q            string `form:"q,omitempty"`            // Limit search to tags that contain the supplied string.
}

// TagsResponse is the response for TagsRequest
type TagsResponse struct {
	Tags []string `json:"tags,omitempty"`
}

// UpdateRequest Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission
type UpdateRequest struct {
	DebtSubCharacteristic      string `form:"debt_sub_characteristic,omitempty"`       // Debt characteristics are no more supported. This parameter is ignored.
	Key                        string `form:"key,omitempty"`                           // Key of the rule to update
	MarkdownDescription        string `form:"markdown_description,omitempty"`          // Rule description (mandatory for custom rule and manual rule)
	MarkdownNote               string `form:"markdown_note,omitempty"`                 // Optional note in markdown format. Use empty value to remove current note. Note is not changed if the parameter is not set.
	Name                       string `form:"name,omitempty"`                          // Rule name (mandatory for custom rule)
	Organization               string `form:"organization,omitempty"`                  // Organization key
	Params                     string `form:"params,omitempty"`                        // Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only when updating a custom rule)
	RemediationFnBaseEffort    string `form:"remediation_fn_base_effort,omitempty"`    // Base effort of the remediation function of the rule
	RemediationFnType          string `form:"remediation_fn_type,omitempty"`           // Type of the remediation function of the rule
	RemediationFyGapMultiplier string `form:"remediation_fy_gap_multiplier,omitempty"` // Gap multiplier of the remediation function of the rule
	Severity                   string `form:"severity,omitempty"`                      // Rule severity (Only when updating a custom rule)
	Status                     string `form:"status,omitempty"`                        // Rule status (Only when updating a custom rule)
	Tags                       string `form:"tags,omitempty"`                          // Optional comma-separated list of tags to set. Use blank value to remove current tags. Tags are not changed if the parameter is not set.
}

// UpdateResponse is the response for UpdateRequest
type UpdateResponse struct {
	Rule struct {
		CreatedAt      string `json:"createdAt,omitempty"`
		DebtOverloaded bool   `json:"debtOverloaded,omitempty"`
		HtmlDesc       string `json:"htmlDesc,omitempty"`
		IsExternal     bool   `json:"isExternal,omitempty"`
		IsTemplate     bool   `json:"isTemplate,omitempty"`
		Key            string `json:"key,omitempty"`
		Lang           string `json:"lang,omitempty"`
		LangName       string `json:"langName,omitempty"`
		MdDesc         string `json:"mdDesc,omitempty"`
		Name           string `json:"name,omitempty"`
		Params         []struct {
			DefaultValue string `json:"defaultValue,omitempty"`
			HtmlDesc     string `json:"htmlDesc,omitempty"`
			Key          string `json:"key,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"params,omitempty"`
		RemFnOverloaded bool     `json:"remFnOverloaded,omitempty"`
		Repo            string   `json:"repo,omitempty"`
		Scope           string   `json:"scope,omitempty"`
		Severity        string   `json:"severity,omitempty"`
		Status          string   `json:"status,omitempty"`
		SysTags         []string `json:"sysTags,omitempty"`
		Tags            []string `json:"tags,omitempty"`
		TemplateKey     string   `json:"templateKey,omitempty"`
		Type            string   `json:"type,omitempty"`
	} `json:"rule,omitempty"`
}
