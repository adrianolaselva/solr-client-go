package solr

type Client struct {
	config *Config
	Document Document
	Collection Collection
}

func NewClient(config *Config) Client {
	collection := NewCollection(config)
	document := NewDocument(config)
	return Client{
		config: config,
		Collection: collection,
		Document: document,
	}
}
