package grpc_clients

import (
	"context"

	userProto "github.com/Teamlix/proto/gen/go/user_service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userProto.UserServiceClient
}

func newUserClient(address string) (*UserClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := userProto.NewUserServiceClient(conn)

	return &UserClient{client: client}, nil
}

func (c UserClient) CheckAccessToken(
	ctx context.Context,
	token string,
) (bool, error) {
	req := userProto.CheckAccessTokenRequest{
		AccessToken: token,
	}
	res, err := c.client.CheckAccessToken(ctx, &req)
	if err != nil {
		return false, err
	}

	v := res.GetResult()
	return v, nil
}

func (c UserClient) GetUserById(
	ctx context.Context,
	id string,
) error {
	return nil
}

func (c UserClient) GetUsersByIds(
	ctx context.Context,
	ids []string,
) error {
	return nil
}
