package dahua_panel

type settingRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
	Session string      `json:"session"`
}

func NewSettingRequest(maintainParams *maintainParams, id int, session string, requestType ...string) *settingRequest {
	rType := configRequestMethodName
	if len(requestType) == 1 {
		rType = requestType[0]
	}

	return &settingRequest{
		Method:  rType,
		Params:  maintainParams,
		ID:      id,
		Session: session,
	}
}
