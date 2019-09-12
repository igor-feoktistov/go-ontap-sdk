package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportPolicyCreate struct {
	Base
	Params struct {
		XMLName xml.Name
		PolicyName   string `xml:"policy-name"`
		ReturnRecord bool   `xml:"return-record,omitempty"`
	}
}

type ExportPolicyCreateResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		result ExportPolicyInfo `xml:"result,omitempty"`
	} `xml:"results"`
}

func (c *Client) ExportPolicyCreateAPI(policyName string, returnRecord bool) (*ExportPolicyCreateResponse, *http.Response, error) {
	if c.ExportPolicyCreate == nil {
		c.ExportPolicyCreate = &ExportPolicyCreate {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportPolicyCreate.Params.XMLName = xml.Name{Local: "export-policy-create"}
	}
	c.ExportPolicyCreate.Base.Name = c.vserver
	c.ExportPolicyCreate.Params.PolicyName = policyName
	c.ExportPolicyCreate.Params.ReturnRecord = returnRecord
	r := ExportPolicyCreateResponse{}
	res, err := c.ExportPolicyCreate.get(c.ExportPolicyCreate, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportPolicyCreateAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
