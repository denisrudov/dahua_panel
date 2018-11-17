package dahua_panel

type LoginResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	ID     int `json:"id"`
	Params struct {
		Authorization string `json:"authorization"`
		Encryption    string `json:"encryption"`
		Mac           string `json:"mac"`
		Random        string `json:"random"`
		Realm         string `json:"realm"`
	} `json:"params"`
	Result  bool   `json:"result"`
	Session string `json:"session"`
}

type SecondLoginResponse struct {
	ID     int `json:"id"`
	Params struct {
		KeepAliveInterval int `json:"keepAliveInterval"`
	} `json:"params"`
	Result  bool   `json:"result"`
	Session string `json:"session"`
}
