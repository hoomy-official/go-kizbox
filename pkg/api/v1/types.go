package v1

type Gateway struct {
	GatewayID    string `json:"gatewayId"`
	Connectivity struct {
		Status          string `json:"status"`
		ProtocolVersion string `json:"protocolVersion"`
	} `json:"connectivity"`
}

type Version struct {
	ProtocolVersion string `json:"protocolVersion"`
}

type Execution struct {
	Owner            string `json:"owner"`
	ID               string `json:"id"`
	ExecutionType    string `json:"executionType"`
	ExecutionSubType string `json:"executionSubType"`
	Description      string `json:"description"`
	StartTime        int    `json:"startTime"`
	ActionGroup      struct {
		Label   string `json:"label"`
		Actions []struct {
			Commands []struct {
				Type       int      `json:"type"`
				Name       string   `json:"name"`
				Parameters []string `json:"parameters"`
			} `json:"commands"`
			DeviceURL string `json:"deviceURL"`
		} `json:"actions"`
	} `json:"actionGroup"`
	State string `json:"state"`
}

type EventRegister struct {
	ID string `json:"id"`
}

type Event struct {
	Name string `json:"name"`
	Raw  map[string]interface{}
}

type Command struct {
	Name       string        `json:"name"`
	Parameters []interface{} `json:"parameters"`
}

type Action struct {
	Commands  []Command `json:"commands"`
	DeviceURL string    `json:"deviceURL"`
}

type Execute struct {
	Label   string   `json:"label"`
	Actions []Action `json:"actions"`
}

type State struct {
	Name  string      `json:"name"`
	Type  int         `json:"type,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Device struct {
	DeviceURL   string  `json:"deviceURL"`
	SubsystemID int     `json:"subsystemId"`
	Available   bool    `json:"available"`
	Type        int     `json:"type"`
	Label       string  `json:"label"`
	Synced      bool    `json:"synced"`
	States      []State `json:"states"`
	Attributes  []struct {
		Name  string      `json:"name"`
		Type  int         `json:"type,omitempty"`
		Value interface{} `json:"value,omitempty"`
	} `json:"attributes"`
	Enabled          bool   `json:"enabled"`
	ControllableName string `json:"controllableName"`
	Definition       struct {
		States []struct {
			Name string `json:"name"`
		} `json:"states"`
		WidgetName string `json:"widgetName"`
		Type       string `json:"type"`
		Commands   []struct {
			CommandName string `json:"commandName"`
			Nparams     int    `json:"nparams"`
		} `json:"commands"`
		UIClass string `json:"uiClass"`
	} `json:"definition"`
	CreationTime int `json:"creationTime"`
}
