package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
	c := ontap.NewClient(
	    "https://myvserver.example.com",
	    &ontap.ClientOptions {
		Version: "1.160",
		BasicAuthUser: "vsadmin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
	    },
	)
	options := &ontap.ExportRuleGetOptions {
			MaxRecords: 1024,
            		Query: &ontap.ExportRuleQuery {
                		ExportRuleInfo: &ontap.ExportRuleInfo {
                			PolicyName: "xpo_vol_go_test01",
					ClientMatch: "192.168.1.5",
                        	},
                	},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.ExportRuleGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, rule := range response.Results.AttributesList.ExportRuleAttributes {
				fmt.Printf("policy: %s, rule-index: %d, client-match: %s, ro-rule: %s, rw-rule: %s\n", rule.PolicyName, rule.RuleIndex, rule.ClientMatch, rule.RoRule, rule.RwRule)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
