package dahua_panel

const (
	configRequestMethodName = "configManager.setConfig"
)

type settingRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
	Session string      `json:"session"`
}

func NewSettingRequest(maintainParams *maintainParams, id int, session string) *settingRequest {

	request := &settingRequest{
		Method:  configRequestMethodName,
		Params:  maintainParams,
		ID:      id,
		Session: session,
	}
	return request
}
