package request

type WebsiteConfig struct {
	Key   string `json:"key" from:"key" example:"about"`
	Value string `json:"value" from:"value" example:"about me"`
}