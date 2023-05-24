package errors

import "encoding/json"

// assert that Derror implement error interface
var _ error = &Derror{}

type Derror struct {
	code       int
	errorType  string
	message    string
	detail     string
	status     string
	traceId    string
	instanceId string
	help       string
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

// get Derror detail
func (d *Derror) Detail() string {
	return d.detail
}

// get Derror status
func (d *Derror) Status() string {
	return d.status
}

// get Derror traceid
func (d *Derror) TraceId() string {
	return d.traceId
}

// get Deror instanceId
func (d *Derror) InstanceId() string {
	return d.instanceId
}

// get Derror help
func (d *Derror) Help() string {
	return d.help
}

type DerrorJsonModel struct {
	dErrError `json:"error"`
}

type dErrError struct {
	code       int    `json:"code"`
	errorType  string `json:"errorType"`
	message    string `json:"message"`
	detail     string `json:"detail"`
	status     string `json:"status"`
	traceId    string `json:"traceId"`
	instanceId string `json:"instanceId"`
	help       string `json:"help"`
}

func (d *Derror) Error() string {
	jsonByte, _ := json.Marshal(DerrorJsonModel{
		dErrError{
			code:       d.code,
			errorType:  d.errorType,
			message:    d.message,
			detail:     d.detail,
			status:     d.status,
			traceId:    d.traceId,
			instanceId: d.instanceId,
			help:       d.help,
		},
	})
	return string(jsonByte)
}

type TerrorOptionsAttribute func(d *Derror)

// return an instance of Derror with the given attribute
func NewDerror(message, errorType, detail, status string, code int, optionAttr ...TerrorOptionsAttribute) *Derror {
	derror := &Derror{
		code:      code,
		message:   message,
		detail:    detail,
		errorType: errorType,
	}
	for _, attr := range optionAttr {
		attr(derror)
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
		code:       dErrJson.code,
		errorType:  dErrJson.errorType,
		message:    dErrJson.message,
		detail:     dErrJson.detail,
		status:     dErrJson.status,
		traceId:    dErrJson.traceId,
		instanceId: dErrJson.instanceId,
		help:       dErrJson.help,
	}, nil
}
