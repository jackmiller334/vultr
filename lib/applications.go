package lib

import (
	"sort"
	"strings"
)

// Application on Vultr
type Application struct {
	ID         int    `json:"APPID,string"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	DeployName string `json:"deploy_name"`
	Surcharge  int    `json:"surcharge"`
}

type applications []Application

func (s applications) Len() int      { return len(s) }
func (s applications) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s applications) Less(i, j int) bool {
	return strings.ToLower(s[i].Name) < strings.ToLower(s[j].Name)
}

// GetPlans returns a list of all available plans on Vultr account
func (c *Client) GetApplications() ([]Application, error) {
	var applicationMap map[string]Application
	if err := c.get(`app/list`, &applicationMap); err != nil {
		return nil, err
	}

	var applicationList []Application
	for _, application := range applicationMap {
		applicationList = append(applicationList, application)
	}
	sort.Sort(applications(appList))
	return applicationList, nil
}
