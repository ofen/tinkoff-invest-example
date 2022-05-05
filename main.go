package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ofen/tinkoff-invest-example/investapi"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const endpoint = "invest-public-api.tinkoff.ru:443"

var token = flag.String("t", "", "tinkoff invest api token")

func main() {
	flag.Parse()

	if *token == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()
	client := NewSandboxServiceClient(*token)
	resp, err := client.GetSandboxAccounts(ctx, &investapi.GetAccountsRequest{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp)
}

func NewSandboxServiceClient(token string) investapi.SandboxServiceClient {
	conn, err := grpc.Dial(
		endpoint,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, "")),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{AccessToken: token})),
	)
	if err != nil {
		panic(err)
	}

	return investapi.NewSandboxServiceClient(conn)
}
