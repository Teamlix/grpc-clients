package grpc_clients

import (
	"errors"
)

type Clients struct {
	User *UserClient
}

func NewClients(
	userServiceAddress string,
) (*Clients, error) {
	userClient, err := newUserClient(userServiceAddress)
	if err != nil {
		return nil, errors.New("error creating user grpc client")
	}

	return &Clients{
		User: userClient,
	}, nil
}
