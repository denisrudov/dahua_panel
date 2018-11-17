package dahua_panel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"strings"
)

type Dahua struct {
	username     string
	password     string
	address      string
	requestCount int
	session      string
	realm        string
}

// Make default Dahua Panel client
func NewDahuaClient(username, password, address string) *Dahua {
	return &Dahua{
		username:     username,
		password:     password,
		address:      address,
		requestCount: 1,
	}
}

// Login to Dahua Panel
func (d *Dahua) Login() (rez bool) {
	rez = false
	if d.makeFirstLogin() {
		rez = d.makeSecondLogin()
	}

	return
}

// Make first Login call for session value
func (d *Dahua) makeFirstLogin() (rez bool) {
	rez = false
	reqObj := newLoginRequest(d.username, "", "", d.requestCount)
	req := getJsonRequest()
	log.Println("Make first login")
	resp, body, errors := d.makeApiCall(req.Post(d.getLoginUrl()).Send(reqObj))

	var loginRes LoginResponse

	err := json.Unmarshal([]byte(body), &loginRes)

	if err == nil && resp.StatusCode == 200 && len(errors) == 0 {
		rez = true
		//fmt.Println(loginRes.session)
		d.setSessionValue(loginRes.Session)
		d.setRealmValue(loginRes.Params.Realm)
		log.Println("First login successful")
	} else {
		log.Println("First login failed")
	}

	return
}

// Make Real Login Call
func (d *Dahua) makeSecondLogin() (rez bool) {
	rez = false
	loginRequest := newSecondLoginRequest(d.username, d.generateUserPassword(), d.session, d.realm, d.requestCount)
	req := getJsonRequest()
	resp, body, errors := d.makeApiCall(req.Post(d.getLoginUrl()).Send(loginRequest))

	loginResponse := new(SecondLoginResponse)

	err := json.Unmarshal([]byte(body), loginResponse)

	if err == nil && resp.StatusCode == 200 && len(errors) == 0 {
		rez = true
		log.Println("Logged successful")
		d.setSessionValue(loginResponse.Session)
	} else {
		log.Println("Logged failed")
	}

	return
}

// Get basic JSON Request
func getJsonRequest() *gorequest.SuperAgent {
	req := gorequest.New()
	req.Set("Content-type", "application/json")
	return req
}

// Generate MD5 password string
func (d *Dahua) generateUserPassword() string {
	sum1 := md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", d.username, d.realm, d.password)))
	md51 := strings.ToUpper(hex.EncodeToString(sum1[:]))
	sum2 := md5.Sum([]byte(fmt.Sprintf("%s:%s:%s", d.username, d.realm, md51)))
	md52 := hex.EncodeToString(sum2[:])
	return strings.ToUpper(md52)
}

// Get Login Url
func (d *Dahua) getLoginUrl() string {
	return fmt.Sprintf("%s%s", d.getBaseUrl(), LoginEndpoint)
}

// Get Base Url
func (d *Dahua) getBaseUrl() string {
	return fmt.Sprintf("http://%s", d.address)
}

// set session value
func (d *Dahua) setSessionValue(s string) (ret *Dahua) {
	ret = d
	d.session = s
	return
}

// make api call
func (d *Dahua) makeApiCall(send *gorequest.SuperAgent) (*http.Response, string, []error) {

	if d.isLogged() == false {
		return nil, "", []error{errors.New("client is not logged")}
	}
	d.requestCount++
	return send.End()
}
func (d *Dahua) setRealmValue(realmValue string) (rez *Dahua) {
	rez = d
	d.realm = realmValue
	return
}

// Update Maintain Params
func (d *Dahua) UpdateMaintainParams(params *maintainParams) (rez error) {
	req := getJsonRequest()
	setRequest := NewSettingRequest(params, d.requestCount, d.session)
	response, _, errs := d.makeApiCall(req.Post(d.getMaintainsUrl()).Send(setRequest))

	if response.StatusCode != 200 || len(errs) > 0 {
		if len(errs) > 0 {
			rez = errs[0]
		} else {
			rez = errors.New("error updating settings")
		}
	}

	return
}

// Get Maintain Url
func (d *Dahua) getMaintainsUrl() string {
	return fmt.Sprintf("%s%s", d.getBaseUrl(), maintainSettingEndpoint)
}

// Check is the client logged or not
func (d *Dahua) isLogged() bool {
	return len(d.session) > 0
}
