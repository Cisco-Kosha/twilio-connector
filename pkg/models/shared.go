package models

type Specification struct {
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	DomainName string `json:"domain_name,omitempty"`
}

type MessageBody struct {
	From string `json:"From"`
	To   string `json:"To"`
	Body string `json:"Body"`
}
