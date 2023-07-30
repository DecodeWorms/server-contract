package errors

import "encoding/json"

// assert that Derror implement error interface
var _ error = &Derror{}

type Derror struct {
	code      int
	errorType string
	message   string
}

// get Derror code
func (d *Derror) Code() int {
	return d.code
}

// get Derror error tzype
func (d *Derror) ErrorType() string {
	return d.errorType
}

// get Derrpr message
func (d *Derror) Message() string {
	return d.message
}

type DerrorJsonModel struct {
	dErrError `json:"error"`
}

type dErrError struct {
	code      int    `json:"code"`
	errorType string `json:"errorType"`
	message   string `json:"message"`
}

func (d *Derror) Error() string {
	jsonByte, _ := json.Marshal(DerrorJsonModel{
		dErrError{
			code:      d.code,
			errorType: d.errorType,
			message:   d.message,
		},
	})
	return string(jsonByte)
}

type TerrorOptionsAttribute func(d *Derror)

// return an instance of Derror with the given attribute
func NewDerror(message, errorType string, code int) *Derror {
	derror := &Derror{
		code:      code,
		message:   message,
		errorType: errorType,
	}
	return derror
}

func NewDerrorFromJsonString(jsonstring string) (*Derror, error) {
	var dErrJson DerrorJsonModel
	err := json.Unmarshal([]byte(jsonstring), &dErrJson)
	if err != nil {
		return nil, err
	}
	return &Derror{
		code:      dErrJson.code,
		errorType: dErrJson.errorType,
		message:   dErrJson.message,
	}, nil
}
