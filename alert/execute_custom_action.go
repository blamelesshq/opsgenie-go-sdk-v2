package alert

import (
	"github.com/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

type ExecuteCustomActionAlertRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Action 			string
	User        	string            `json:"user,omitempty"`
	Source      	string            `json:"source,omitempty"`
	Note        	string            `json:"note,omitempty"`
}

func (r ExecuteCustomActionAlertRequest) Validate() error {
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	if r.Action == "" {
		return errors.New("Action can not be empty")
	}
	return nil
}

func (r ExecuteCustomActionAlertRequest) Endpoint() string {
	if r.IdentifierType == TINYID {
		return "/v2/alerts/" + r.IdentifierValue + "/actions/"+r.Action+"?identifierType=tiny"
	}else if r.IdentifierType == ALIAS {
		return "/v2/alerts/" + r.IdentifierValue + "/actions/"+r.Action+"?identifierType=alias"
	}
	return "/v2/alerts/" + r.IdentifierValue + "/actions/"+r.Action+"?identifierType=id"

}

func (r ExecuteCustomActionAlertRequest) Method() string {
	return "POST"
}