package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportRuleCreate struct {
	Base
	Params struct {
		XMLName xml.Name
		ExportRuleCreateOptions
	}
}

type ExportRuleCreateOptions struct {
	AnonymousUserId           string    `xml:"anonymous-user-id,omitempty"`
	ClientMatch               string    `xml:"client-match,omitempty"`
	ExportChownMode           string    `xml:"export-chown-mode,omitempty"`
	ExportNtfsUnixSecurityOps string    `xml:"export-ntfs-unix-security-ops,omitempty"`
	IsAllowDevIsEnabled       bool      `xml:"is-allow-dev-is-enabled,omitempty"`
	IsAllowSetUidEnabled      bool      `xml:"is-allow-set-uid-enabled,omitempty"`
	PolicyName                string    `xml:"policy-name,omitempty"`
	Protocol                  *[]string `xml:"protocol>access-protocol,omitempty"`
	RoRule                    *[]string `xml:"ro-rule>security-flavor,omitempty"`
	RuleIndex                 int       `xml:"rule-index,omitempty"`
	RwRule                    *[]string `xml:"rw-rule>security-flavor,omitempty"`
	SuperUserSecurity         *[]string `xml:"super-user-security>security-flavor,omitempty"`
	Vserver                   string    `xml:"vserver-name,omitempty"`
}

func (c *Client) ExportRuleCreateAPI(options *ExportRuleCreateOptions) (*SingleResultResponse, *http.Response, error) {
	if c.ExportRuleCreate == nil {
		c.ExportRuleCreate = &ExportRuleCreate {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportRuleCreate.Params.XMLName = xml.Name{Local: "export-rule-create"}
	}
	c.ExportRuleCreate.Base.Name = c.vserver
	c.ExportRuleCreate.Params.ExportRuleCreateOptions = *options
	r := SingleResultResponse{}
	res, err := c.ExportRuleCreate.get(c.ExportRuleCreate, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportRuleCreateAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
