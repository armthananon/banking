package gapi

import (
	"fmt"

	db "github.com/armthananon/banking/db/sqlc"
	"github.com/armthananon/banking/pb"
	"github.com/armthananon/banking/token"
	"github.com/armthananon/banking/util"
)

// Server serves gRPC requests for our banking services.
type Server struct {
	pb.UnimplementedBankingServiceServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
