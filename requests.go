package dahua_panel

type LoginRequest struct {
	Method  string `json:"method"`
	Params  `json:"params"`
	ID      int         `json:"id"`
	Session interface{} `json:"session"`
}

type SecondLoginRequest struct {
	Method  string           `json:"method"`
	Params  AdditionalParams `json:"params"`
	ID      int              `json:"id"`
	Session interface{}      `json:"session"`
}

type Params struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	ClientType string `json:"clientType"`
}

type AdditionalParams struct {
	Params
	Realm         string `json:"realm"`
	Random        string `json:"random"`
	PasswordType  string `json:"passwordType"`
	AuthorityType string `json:"authorityType"`
}

func newLoginRequest(username, password, session string, id int) *LoginRequest {

	var innerSession interface{}

	if len(session) < 1 {
		innerSession = 0
	} else {
		innerSession = session
	}

	return &LoginRequest{
		Method:  loginRequestMethodName,
		Params:  Params{UserName: username, Password: password, ClientType: ClientType},
		ID:      id,
		Session: innerSession,
	}
}

func newSecondLoginRequest(username, password, session, realm string, id int) *SecondLoginRequest {
	return &SecondLoginRequest{
		Params: AdditionalParams{
			Params:        Params{UserName: username, Password: password, ClientType: ClientType},
			Realm:         realm,
			AuthorityType: AuthorityType,
			PasswordType:  PasswordType,
			Random:        realm,
		},
		Session: session,
		ID:      id,
		Method:  loginRequestMethodName,
	}
}
