package lib

// Application on Vultr
type Application struct {
    ID int `json:"APPID,string"`
    Name string `json:"name"`
    ShortName string `json:"short_name"`
    DeployName string `json:"deploy_name"`
    Surcharge int `json:"surcharge"`
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
	return applicationList, nil
}