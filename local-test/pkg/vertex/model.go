package vertex

type SearchRequest struct {
	Query               string `json:"query"`
	PageSize            int    `json:"pageSize"`
	ContentSearchSpec   `json:"contentSearchSpec"`
	QueryExpansionSpec  `json:"queryExpansionSpec"`
	SpellCorrectionSpec `json:"spellCorrectionSpec"`
}

type ContentSearchSpec struct {
	SnippetSpec SnippetSpec `json:"snippetSpec"`
	SummarySpec SummarySpec `json:"summarySpec"`
}

type SnippetSpec struct {
	ReturnSnippet bool `json:"returnSnippet"`
}

type SummarySpec struct {
	SummaryResultCount           int  `json:"summaryResultCount"`
	IncludeCitations             bool `json:"includeCitations"`
	IgnoreAdversarialQuery       bool `json:"ignoreAdversarialQuery"`
	IgnoreNonSummarySeekingQuery bool `json:"ignoreNonSummarySeekingQuery"`
	ModelPromptSpec              `json:"modelPromptSpec"`
	ModelSpec                    `json:"modelSpec"`
}

type ModelPromptSpec struct {
	Preamble string `json:"preamble"`
}

type ModelSpec struct {
	Version string `json:"version"`
}

type QueryExpansionSpec struct {
	Condition string `json:"condition"`
}

type SpellCorrectionSpec struct {
	Mode string `json:"mode"`
}