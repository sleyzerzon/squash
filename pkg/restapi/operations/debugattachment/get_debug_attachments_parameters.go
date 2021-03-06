// Code generated by go-swagger; DO NOT EDIT.

package debugattachment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDebugAttachmentsParams creates a new GetDebugAttachmentsParams object
// with the default values initialized.
func NewGetDebugAttachmentsParams() GetDebugAttachmentsParams {
	var ()
	return GetDebugAttachmentsParams{}
}

// GetDebugAttachmentsParams contains all the bound params for the get debug attachments operation
// typically these are obtained from a http.Request
//
// swagger:parameters getDebugAttachments
type GetDebugAttachmentsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: header
	*/
	XTimeout *float64
	/*Only get a subset of debugattachments
	  In: query
	*/
	Names []string
	/*filter by node that the debugattachment is assigned to
	  In: query
	*/
	Node *string
	/*filter by the state of debugattachment is assigned to
	  In: query
	*/
	State *string
	/*filter by any of the states of the debugattachment (for example, attached and error)
	  In: query
	*/
	States []string
	/*wait until there's something to return instead of returning an empty list
	  In: query
	*/
	Wait *bool
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetDebugAttachmentsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXTimeout(r.Header[http.CanonicalHeaderKey("X-Timeout")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qNames, qhkNames, _ := qs.GetOK("names")
	if err := o.bindNames(qNames, qhkNames, route.Formats); err != nil {
		res = append(res, err)
	}

	qNode, qhkNode, _ := qs.GetOK("node")
	if err := o.bindNode(qNode, qhkNode, route.Formats); err != nil {
		res = append(res, err)
	}

	qState, qhkState, _ := qs.GetOK("state")
	if err := o.bindState(qState, qhkState, route.Formats); err != nil {
		res = append(res, err)
	}

	qStates, qhkStates, _ := qs.GetOK("states")
	if err := o.bindStates(qStates, qhkStates, route.Formats); err != nil {
		res = append(res, err)
	}

	qWait, qhkWait, _ := qs.GetOK("wait")
	if err := o.bindWait(qWait, qhkWait, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDebugAttachmentsParams) bindXTimeout(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("X-Timeout", "header", "float64", raw)
	}
	o.XTimeout = &value

	return nil
}

func (o *GetDebugAttachmentsParams) bindNames(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvNames string
	if len(rawData) > 0 {
		qvNames = rawData[len(rawData)-1]
	}

	namesIC := swag.SplitByFormat(qvNames, "")

	if len(namesIC) == 0 {
		return nil
	}

	var namesIR []string
	for _, namesIV := range namesIC {
		namesI := namesIV

		namesIR = append(namesIR, namesI)
	}

	o.Names = namesIR

	return nil
}

func (o *GetDebugAttachmentsParams) bindNode(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Node = &raw

	return nil
}

func (o *GetDebugAttachmentsParams) bindState(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.State = &raw

	return nil
}

func (o *GetDebugAttachmentsParams) bindStates(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvStates string
	if len(rawData) > 0 {
		qvStates = rawData[len(rawData)-1]
	}

	statesIC := swag.SplitByFormat(qvStates, "")

	if len(statesIC) == 0 {
		return nil
	}

	var statesIR []string
	for _, statesIV := range statesIC {
		statesI := statesIV

		statesIR = append(statesIR, statesI)
	}

	o.States = statesIR

	return nil
}

func (o *GetDebugAttachmentsParams) bindWait(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("wait", "query", "bool", raw)
	}
	o.Wait = &value

	return nil
}
