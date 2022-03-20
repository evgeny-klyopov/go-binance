package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListLiquidSwapProductsService https://binance-docs.github.io/apidocs/spot/en/#get-fixed-and-activity-project-list-user_data
type ListLiquidSwapProductsService struct {
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
func (s *ListLiquidSwapProductsService) Asset(asset string) *ListLiquidSwapProductsService {
	s.asset = asset
	return s
}

// Type set project type ("ACTIVITY", "CUSTOMIZED_FIXED")
func (s *ListLiquidSwapProductsService) Type(projectType string) *ListLiquidSwapProductsService {
	s.projectType = projectType
	return s
}

// IsSortAsc default "true"
func (s *ListLiquidSwapProductsService) IsSortAsc(isSortAsc bool) *ListLiquidSwapProductsService {
	s.isSortAsc = isSortAsc
	return s
}

// Status ("ALL", "SUBSCRIBABLE", "UNSUBSCRIBABLE") - default "ALL"
func (s *ListLiquidSwapProductsService) Status(status string) *ListLiquidSwapProductsService {
	s.status = status
	return s
}

// SortBy ("START_TIME", "LOT_SIZE", "INTEREST_RATE", "DURATION") - default "START_TIME"
func (s *ListLiquidSwapProductsService) SortBy(sortBy string) *ListLiquidSwapProductsService {
	s.sortBy = sortBy
	return s
}

// Current Currently querying page. Start from 1. Default:1
func (s *ListLiquidSwapProductsService) Current(current int64) *ListLiquidSwapProductsService {
	s.current = current
	return s
}

// Size Default:10, Max:100
func (s *ListLiquidSwapProductsService) Size(size int64) *ListLiquidSwapProductsService {
	s.size = size
	return s
}

// Do send request
func (s *ListLiquidSwapProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*ListLiquidSwapProducts, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidityOps",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	var res []*ListLiquidSwapProducts
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// SavingsFixedProduct define a fixed product (Savings)
type ListLiquidSwapProducts struct {
	OperationId int    `json:"operationId"`
	PoolId      int    `json:"poolId"`
	PoolName    string `json:"poolName"`
	Operation   string `json:"operation"`
	Status      int    `json:"status"`
	UpdateTime  int64  `json:"updateTime"`
	ShareAmount string `json:"shareAmount"`
}


type LiquidSwapProductsService struct {
	c           *Client
	poolId       string
	projectType string
	status      string
	isSortAsc   bool
	sortBy      string
	current     int64
	size        int64
}

// Asset desired asset
func (s *LiquidSwapProductsService) PoolId(poolId string) *LiquidSwapProductsService {
	s.poolId = poolId
	return s
}

// Type set project type ("ACTIVITY", "CUSTOMIZED_FIXED")
func (s *LiquidSwapProductsService) Type(projectType string) *LiquidSwapProductsService {
	s.projectType = projectType
	return s
}

// IsSortAsc default "true"
func (s *LiquidSwapProductsService) IsSortAsc(isSortAsc bool) *LiquidSwapProductsService {
	s.isSortAsc = isSortAsc
	return s
}

// Status ("ALL", "SUBSCRIBABLE", "UNSUBSCRIBABLE") - default "ALL"
func (s *LiquidSwapProductsService) Status(status string) *LiquidSwapProductsService {
	s.status = status
	return s
}

// SortBy ("START_TIME", "LOT_SIZE", "INTEREST_RATE", "DURATION") - default "START_TIME"
func (s *LiquidSwapProductsService) SortBy(sortBy string) *LiquidSwapProductsService {
	s.sortBy = sortBy
	return s
}

// Current Currently querying page. Start from 1. Default:1
func (s *LiquidSwapProductsService) Current(current int64) *LiquidSwapProductsService {
	s.current = current
	return s
}

// Size Default:10, Max:100
func (s *LiquidSwapProductsService) Size(size int64) *LiquidSwapProductsService {
	s.size = size
	return s
}

// Do send request
func (s *LiquidSwapProductsService) Do(ctx context.Context, opts ...RequestOption) ([]*LiquidSwapProducts, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/bswap/liquidity",
		secType:  secTypeSigned,
	}

	m := params{}
	m["poolId"] = s.poolId
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)


	if err != nil {
		return nil, err
	}
	var res []*LiquidSwapProducts
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return res, nil
}

type LiquidityShare struct {
	ShareAmount     string `json:"shareAmount"`
	SharePercentage string `json:"sharePercentage"`
	Asset           map[string]string `json:"asset"`
}

// SavingsFixedProduct define a fixed product (Savings)
type LiquidSwapProducts struct {
	PoolId     int    `json:"poolId"`
	PoolName   string `json:"poolName"`
	UpdateTime int64  `json:"updateTime"`
	Liquidity  map[string]string `json:"liquidity"`
	Share LiquidityShare `json:"share"`
}