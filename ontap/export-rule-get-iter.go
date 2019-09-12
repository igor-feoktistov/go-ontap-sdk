package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type ExportRuleGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*ExportRuleGetOptions
	}
}

type ExportRuleGetOptions struct {
	DesiredAttributes *ExportRuleQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int              `xml:"max-records,omitempty"`
	Query             *ExportRuleQuery `xml:"query,omitempty"`
	Tag               string           `xml:"tag,omitempty"`
}

type ExportRuleQuery struct {
	ExportRuleInfo *ExportRuleInfo `xml:"export-rule-info,omitempty"`
}

type ExportRuleInfo struct {
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

type ExportRuleGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			ExportRuleAttributes []ExportRuleInfo `xml:"export-rule-info,omitempty"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) ExportRuleGetAPI(options *ExportRuleGetOptions) (*ExportRuleGetResponse, *http.Response, error) {
	if c.ExportRuleGetIter == nil {
		c.ExportRuleGetIter = &ExportRuleGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.ExportRuleGetIter.Params.XMLName = xml.Name{Local: "export-rule-get-iter"}
	}
	c.ExportRuleGetIter.Base.Name = c.vserver
	c.ExportRuleGetIter.Params.ExportRuleGetOptions = options
	r := ExportRuleGetResponse{}
	res, err := c.ExportRuleGetIter.get(c.ExportRuleGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(ExportRuleGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) ExportRuleGetIterAPI(options *ExportRuleGetOptions) (responses []*ExportRuleGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.ExportRuleGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responses = append(responses, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
