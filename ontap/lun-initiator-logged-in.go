package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type LunInitiatorLoggedIn struct {
	Base
	Params struct {
		XMLName    xml.Name
		Initiator  string `xml:"initiator"`
	}
}

type LunInitiatorLoggedInResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		IsLoggedIn bool `xml:"is-logged-in"`
	} `xml:"results"`
}

func (c *Client) LunInitiatorLoggedInAPI(initiatorName string) (*LunInitiatorLoggedInResponse, *http.Response, error) {
	if c.LunInitiatorLoggedIn == nil {
		c.LunInitiatorLoggedIn = &LunInitiatorLoggedIn {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.LunInitiatorLoggedIn.Params.XMLName = xml.Name{Local: "lun-initiator-logged-in"}
	}
	c.LunInitiatorLoggedIn.Base.Name = c.vserver
	c.LunInitiatorLoggedIn.Params.Initiator = initiatorName
	r := LunInitiatorLoggedInResponse{}
	res, err := c.LunInitiatorLoggedIn.get(c.LunInitiatorLoggedIn, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(LunInitiatorLoggedInAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}
