package ytlmetadata

type metadataRequest struct {
	Context context `json:"context"`
	VideoID string  `json:"videoId"`
}

type context struct {
	Client client `json:"client"`
}

type client struct {
	Hl            string `json:"hl"`
	ClientName    string `json:"clientName"`
	ClientVersion string `json:"clientVersion"`
}

type metadataResponse struct {
	ResponseContext responseContext `json:"responseContext"`
	Actions         []actions       `json:"actions"`
	Continuation    continuation    `json:"continuation"`
}

type params struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type serviceTrackingParams struct {
	Service string   `json:"service"`
	Params  []params `json:"params"`
}

type webResponseContextExtensionData struct {
	HasDecorated bool `json:"hasDecorated"`
}

type responseContext struct {
	VisitorData                     string                          `json:"visitorData"`
	ServiceTrackingParams           []serviceTrackingParams         `json:"serviceTrackingParams"`
	WebResponseContextExtensionData webResponseContextExtensionData `json:"webResponseContextExtensionData"`
}

type defaultText struct {
	SimpleText string `json:"simpleText"`
}

type toggledText struct {
	SimpleText string `json:"simpleText"`
}

type updateToggleButtonTextAction struct {
	DefaultText defaultText `json:"defaultText"`
	ToggledText toggledText `json:"toggledText"`
	ButtonID    string      `json:"buttonId"`
}

type dateText struct {
	SimpleText string `json:"simpleText"`
}

type updateDateTextAction struct {
	DateText dateText `json:"dateText"`
}

type title struct {
	SimpleText string `json:"simpleText"`
}

type updateTitleAction struct {
	Title title `json:"title"`
}

type webCommandMetadata struct {
	URL         string `json:"url"`
	WebPageType string `json:"webPageType"`
	RootVe      int    `json:"rootVe"`
}

type commandMetadata struct {
	WebCommandMetadata webCommandMetadata `json:"webCommandMetadata"`
}

type searchEndpoint struct {
	Query string `json:"query"`
}

type navigationEndpoint struct {
	CommandMetadata     commandMetadata `json:"commandMetadata"`
	SearchEndpoint      searchEndpoint  `json:"searchEndpoint"`
	WatchEndpoint       watchEndpoint   `json:"watchEndpoint"`
	ClickTrackingParams string          `json:"clickTrackingParams"`
	URLEndpoint         urlEndpoint     `json:"urlEndpoint"`
}

type Visibility struct {
	Types string `json:"types"`
}

type LoggingDirectives struct {
	Visibility Visibility `json:"visibility"`
}

type watchEndpoint struct {
	VideoID string `json:"videoId"`
}

type urlEndpoint struct {
	URL      string `json:"url"`
	Nofollow bool   `json:"nofollow"`
}

type runs struct {
	Text               string             `json:"text"`
	NavigationEndpoint navigationEndpoint `json:"navigationEndpoint,omitempty"`
	LoggingDirectives  LoggingDirectives  `json:"loggingDirectives,omitempty"`
}

type description struct {
	Runs []runs `json:"runs"`
}

type updateDescriptionAction struct {
	Description description `json:"description"`
}

type actions struct {
	UpdateToggleButtonTextAction updateToggleButtonTextAction `json:"updateToggleButtonTextAction,omitempty"`
	UpdateDateTextAction         updateDateTextAction         `json:"updateDateTextAction,omitempty"`
	UpdateTitleAction            updateTitleAction            `json:"updateTitleAction,omitempty"`
	UpdateDescriptionAction      updateDescriptionAction      `json:"updateDescriptionAction,omitempty"`
}

type timedContinuationData struct {
	TimeoutMs    int    `json:"timeoutMs"`
	Continuation string `json:"continuation"`
}

type continuation struct {
	TimedContinuationData timedContinuationData `json:"timedContinuationData"`
}
