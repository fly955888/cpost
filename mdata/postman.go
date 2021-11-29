package mdata

type Postman struct {
	QueryType          string       `json:"postType"`
	QueryUrl           string       `json:"postUrl"`
	QueryParam         []QueryParam `json:"queryParam"`
	QueryAuthorization QueryAuthorization
	QueryHeaders       QueryHeaders
}

type QueryParam struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type QueryAuthorization struct {
	Type []string `json:"type"`
}

type QueryHeaders struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type QueryBody struct {
	FormData           FormData
	XWWWFormUrlencoded XWWWFormUrlencoded
	Raw                Raw
}

type FormData struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type XWWWFormUrlencoded struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type Raw struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
