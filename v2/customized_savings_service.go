package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListCustomizedAndFixedProductsService https://binance-docs.github.io/apidocs/spot/en/#get-fixed-and-activity-project-list-user_data
type ListCustomizedAndFixedProductsService struct {
	c           *Client
	asset       string
	projectType string
	status      string
	isSortAsc   bool
	sortBy      string
	current     int64
	size        int64
}

// Asset desired asset
func (s *ListCustomizedAndFixedProductsService) Asset(asset string) *ListCustomizedAndFixedProductsService {
	s.asset = asset
	return s
}

// Type set project type ("ACTIVITY", "CUSTOMIZED_FIXED")
func (s *ListCustomizedAndFixedProductsService) Type(projectType string) *ListCustomizedAndFixedProductsService {
	s.projectType = projectType
	return s
}

// IsSortAsc default "true"
func (s *ListCustomizedAndFixedProductsService) IsSortAsc(isSortAsc bool) *ListCustomizedAndFixedProductsService {
	s.isSortAsc = isSortAsc
	return s
}

// Status ("ALL", "SUBSCRIBABLE", "UNSUBSCRIBABLE") - default "ALL"
func (s *ListCustomizedAndFixedProductsService) Status(status string) *ListCustomizedAndFixedProductsService {
	s.status = status
	return s
}

// SortBy ("START_TIME", "LOT_SIZE", "INTEREST_RATE", "DURATION") - default "START_TIME"
func (s *ListCustomizedAndFixedProductsService) SortBy(sortBy string) *ListCustomizedAndFixedProductsService {
	s.sortBy = sortBy
	return s
}

// Current Currently querying page. Start from 1. Default:1
func (s *ListCustomizedAndFixedProductsService) Current(current int64) *ListCustomizedAndFixedProductsService {
	s.current = current
	return s
}

// Size Default:10, Max:100
func (s *ListCustomizedAndFixedProductsService) Size(size int64) *ListCustomizedAndFixedProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListCustomizedAndFixedProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*CustomizedAndFixedProductsService, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/lending/project/position/list",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*CustomizedAndFixedProductsService
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFixedProduct define a fixed product (Savings)
type CustomizedAndFixedProductsService struct {
	Asset           string `json:"asset"`
	CanTransfer     bool   `json:"canTransfer"`
	CreateTimestamp int64  `json:"createTimestamp"`
	Duration        int    `json:"duration"`
	EndTime         int64  `json:"endTime"`
	Interest        string `json:"interest"`
	InterestRate    string `json:"interestRate"`
	Lot             int    `json:"lot"`
	PositionId      int    `json:"positionId"`
	Principal       string `json:"principal"`
	ProjectId       string `json:"projectId"`
	ProjectName     string `json:"projectName"`
	PurchaseTime    int64  `json:"purchaseTime"`
	RedeemDate      string `json:"redeemDate"`
	StartTime       int64  `json:"startTime"`
	Status          string `json:"status"`
	Type            string `json:"type"`
}
