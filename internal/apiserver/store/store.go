package store

var client Factory

// Factory 仓库层
type Factory interface {
	Users() UserStore
	Secrets() SecretStore
	Close() error
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
