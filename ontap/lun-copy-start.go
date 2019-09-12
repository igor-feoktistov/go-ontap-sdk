package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunCopyStart struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunCopyStartOptions
	}
}

type LunCopyStartOptions struct {
	MaxThroughput int            `xml:"max-throughput,omitempty"`
	Paths         *[]LunPathPair `xml:"paths>lun-path-pair"`
	PromoteEarly  bool           `xml:"promote-early,omitempty"`
	ReferencePath string         `xml:"reference-path,omitempty"`
	SourceVserver string         `xml:"source-vserver,omitempty"`
}

type LunPathPair struct {
	DestinationPath string `xml:"destination-path"`
	SourcePath      string `xml:"source-path"`
}

type LunCopyStartResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		JobUuid string `xml:"job-uuid"`
	} `xml:"results"`
}

func (c *Client) LunCopyStartAPI(options *LunCopyStartOptions) (*LunCopyStartResponse, *http.Response, error) {
	if c.LunCopyStart == nil {
		c.LunCopyStart = &LunCopyStart {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunCopyStart.Params.XMLName = xml.Name{Local: "lun-copy-start"}
	}
	c.LunCopyStart.Base.Name = c.vserver
	c.LunCopyStart.Params.LunCopyStartOptions = *options
	r := LunCopyStartResponse{}
	res, err := c.LunCopyStart.get(c.LunCopyStart, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunCopyStartAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
