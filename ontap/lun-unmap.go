package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunUnmap struct {
	Base
	Params struct {
		XMLName    xml.Name
		LunUnmapOptions
	}
}

type LunUnmapOptions struct {
	InitiatorGroup string `xml:"initiator-group"`
	Path           string `xml:"path"`
}

func (c *Client) LunUnmapAPI(options *LunUnmapOptions) (*SingleResultResponse, *http.Response, error) {
	if c.LunUnmap == nil {
		c.LunUnmap = &LunUnmap {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunUnmap.Params.XMLName = xml.Name{Local: "lun-unmap"}
	}
	c.LunUnmap.Base.Name = c.vserver
	c.LunUnmap.Params.LunUnmapOptions = *options
	r := SingleResultResponse{}
	res, err := c.LunUnmap.get(c.LunUnmap, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunUnmapAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
