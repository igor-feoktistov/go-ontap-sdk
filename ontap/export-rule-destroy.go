package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportRuleDestroy struct {
	Base
	Params struct {
		XMLName xml.Name
		PolicyName string `xml:"policy-name"`
		RuleIndex  int    `xml:"rule-index"`
	}
}

func (c *Client) ExportRuleDestroyAPI(policyName string, ruleIndex int) (*SingleResultResponse, *http.Response, error) {
	if c.ExportRuleDestroy == nil {
		c.ExportRuleDestroy = &ExportRuleDestroy {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportRuleDestroy.Params.XMLName = xml.Name{Local: "export-rule-destroy"}
	}
	c.ExportRuleDestroy.Base.Name = c.vserver
	c.ExportRuleDestroy.Params.PolicyName = policyName
	c.ExportRuleDestroy.Params.RuleIndex = ruleIndex
	r := SingleResultResponse{}
	res, err := c.ExportRuleDestroy.get(c.ExportRuleDestroy, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportRuleDestroyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
