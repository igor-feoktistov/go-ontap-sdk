package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportRuleModify struct {
	Base
	Params struct {
		XMLName xml.Name
		ExportRuleModifyOptions
	}
}

type ExportRuleModifyOptions struct {
	AnonymousUserId           string    `xml:"anonymous-user-id,omitempty"`
	ClientMatch               string    `xml:"client-match,omitempty"`
	ExportChownMode           string    `xml:"export-chown-mode,omitempty"`
	ExportNtfsUnixSecurityOps string    `xml:"export-ntfs-unix-security-ops,omitempty"`
	IsAllowDevIsEnabled       bool      `xml:"is-allow-dev-is-enabled,omitempty"`
	IsAllowSetUidEnabled      bool      `xml:"is-allow-set-uid-enabled,omitempty"`
	PolicyName                string    `xml:"policy-name"`
	Protocol                  *[]string `xml:"protocol>access-protocol,omitempty"`
	RoRule                    *[]string `xml:"ro-rule>security-flavor,omitempty"`
	RuleIndex                 int       `xml:"rule-index"`
	RwRule                    *[]string `xml:"rw-rule>security-flavor,omitempty"`
	SuperUserSecurity         *[]string `xml:"super-user-security>security-flavor,omitempty"`
}

func (c *Client) ExportRuleModifyAPI(options *ExportRuleModifyOptions) (*SingleResultResponse, *http.Response, error) {
	if c.ExportRuleModify == nil {
		c.ExportRuleModify = &ExportRuleModify {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportRuleModify.Params.XMLName = xml.Name{Local: "export-rule-modify"}
	}
	c.ExportRuleModify.Base.Name = c.vserver
	c.ExportRuleModify.Params.ExportRuleModifyOptions = *options
	r := SingleResultResponse{}
	res, err := c.ExportRuleModify.get(c.ExportRuleModify, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportRuleModifyAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
