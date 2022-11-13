package entities

type Client struct {
	IpClient  string
	UserAgent string
}

// Builder Object for Client
type ClientBuilder struct {
	ipClient  string
	userAgent string
}

// Constructor for ClientBuilder
func NewClientBuilder() *ClientBuilder {
	o := new(ClientBuilder)
	return o
}

// Build Method which creates Client
func (b *ClientBuilder) Build() *Client {
	o := new(Client)
	o.IpClient = b.ipClient
	o.UserAgent = b.userAgent
	return o
}

// Setter method for the field ipClient of type string in the object ClientBuilder
func (c *ClientBuilder) SetIpClient(ipClient string) {
	c.ipClient = ipClient
}

// Setter method for the field userAgent of type string in the object ClientBuilder
func (c *ClientBuilder) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}
