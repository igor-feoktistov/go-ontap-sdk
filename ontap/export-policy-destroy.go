package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportPolicyDestroy struct {
	Base
	Params struct {
		XMLName xml.Name
		PolicyName string `xml:"policy-name"`
	}
}

func (c *Client) ExportPolicyDestroyAPI(policyName string) (*SingleResultResponse, *http.Response, error) {
	if c.ExportPolicyDestroy == nil {
		c.ExportPolicyDestroy = &ExportPolicyDestroy {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportPolicyDestroy.Params.XMLName = xml.Name{Local: "export-policy-destroy"}
	}
	c.ExportPolicyDestroy.Base.Name = c.vserver
	c.ExportPolicyDestroy.Params.PolicyName = policyName
	r := SingleResultResponse{}
	res, err := c.ExportPolicyDestroy.get(c.ExportPolicyDestroy, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportPolicyDestroyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
