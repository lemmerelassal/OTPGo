package backend

import (
	"context"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/sirupsen/logrus"

	//xds "x-cite.io/brain/pkg/iot/driver/datastore"
	backendp "otpgo/pkg/proto/backend"
	//device "x-cite.io/brain/pkg/iot/proto/devicemanagement"
)

// Server Definition
type Server struct {
	logger *logrus.Entry
}

// Constructor
func New() *Server {
	return &Server{
		logger: logrus.NewEntry(logrus.StandardLogger()),
	}
}

func (s *Server) Login(ctx context.Context, in *backendp.LoginRequest) (*backendp.ResponseStatus, error) {
	var response backendp.ResponseStatus

	response.Status = "Works!"

	return &response, nil
}

func (s *Server) Logout(ctx context.Context, in *backendp.LogoutRequest) (*backendp.ResponseStatus, error) {
	var response backendp.ResponseStatus

	response.Status = "Also works!"

	return &response, nil
}

func (s *Server) UpdateProfile(ctx context.Context, in *backendp.Context) (*backendp.ResponseStatus, error) {

	var response backendp.ResponseStatus

	response.Status = "Yep, this also works!"

	return &response, nil
}

func (s *Server) Signup(ctx context.Context, in *backendp.Context) (*backendp.ResponseStatus, error) {
	var response backendp.ResponseStatus

	response.Status = "Same!"

	return &response, nil
}

func (s *Server) GetTicker(ctx context.Context, in *backendp.GetTickerRequest) (*backendp.GetTickerResponse, error) {
	var response backendp.GetTickerResponse
	var t timestamp.Timestamp
	t.Seconds = 1000000
	t.Nanos = 0

	response.High = 100
	response.Low = 10
	response.Last = 100
	response.TimeOpen = &t
	response.TimeClose = &t
	response.Volume = 1000000

	return &response, nil
}

func (s *Server) GetOrderBook(ctx context.Context, in *backendp.GetOrderBookRequest) (*backendp.GetOrderBookResponse, error) {
	var response backendp.GetOrderBookResponse
	response.Asks = make([]*backendp.OrderBookEntry, 2)
	response.Bids = make([]*backendp.OrderBookEntry, 2)
	var temp [4]backendp.OrderBookEntry
	response.Asks[0] = &temp[0]
	response.Asks[1] = &temp[1]
	response.Bids[0] = &temp[2]
	response.Bids[1] = &temp[3]

	response.Asks[0].Amount = 10
	response.Asks[0].Price = 10
	response.Asks[0].Total = 100
	response.Asks[1].Amount = 20
	response.Asks[1].Price = 20
	response.Asks[1].Total = 200
	response.Bids[0].Amount = 10
	response.Bids[0].Price = 10
	response.Bids[0].Total = 100
	response.Bids[1].Amount = 20
	response.Bids[1].Price = 20
	response.Bids[1].Total = 200

	return &response, nil
}

func (s *Server) GetBalance(ctx context.Context, in *backendp.GetBalanceRequest) (*backendp.GetBalanceResponse, error) {
	var response backendp.GetBalanceResponse

	return &response, nil
}

func (s *Server) CancelOrder(ctx context.Context, in *backendp.CancelOrderRequest) (*backendp.ResponseStatus, error) {
	var response backendp.ResponseStatus

	return &response, nil
}

func (s *Server) GetOpenOrders(ctx context.Context, in *backendp.GetOpenOrdersRequest) (*backendp.GetOpenOrdersResponse, error) {
	var response backendp.GetOpenOrdersResponse

	return &response, nil
}

func (s *Server) PlaceOrder(ctx context.Context, in *backendp.PlaceOrderRequest) (*backendp.ResponseStatus, error) {
	var response backendp.ResponseStatus

	return &response, nil
}

func (s *Server) GetProfile(ctx context.Context, in *backendp.GetProfileRequest) (*backendp.Context, error) {
	var response backendp.Context

	return &response, nil
}

func (s *Server) GetTransactionHistory(ctx context.Context, in *backendp.GetTransactionHistoryRequest) (*backendp.GetTransactionHistoryResponse, error) {
	var response backendp.GetTransactionHistoryResponse

	return &response, nil
}

func (s *Server) GetHistory(ctx context.Context, in *backendp.GetHistoryRequest) (*backendp.GetHistoryResponse, error) {
	var response backendp.GetHistoryResponse

	return &response, nil
}
