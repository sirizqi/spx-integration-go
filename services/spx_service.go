package services

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"spx-integration/config"
	"spx-integration/models"
	"spx-integration/utils"
)

func appIDu64() uint64 {
	v, _ := strconv.ParseUint(config.Cfg.AppID, 10, 64)
	return v
}
func userIDu64() uint64 {
	v, _ := strconv.ParseUint(config.Cfg.UserID, 10, 64)
	return v
}

func signedHeaders(payload []byte, ts int64, rn int64) (map[string]string, error) {
	sign, err := utils.GenerateCheckSign(appIDu64(), config.Cfg.AppSecret, ts, rn, payload)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"Content-Type": "application/json",
		"app-id":       config.Cfg.AppID,
		"check-sign":   sign,
		"timestamp":    strconv.FormatInt(ts, 10),
		"random-num":   strconv.FormatInt(rn, 10),
	}, nil
}

func doPOST(ctx context.Context, path string, reqBody any) ([]byte, int, error) {
	ts := time.Now().Unix()
	rn := rand.New(rand.NewSource(ts)).Int63()
	body, _ := json.Marshal(reqBody)
	h, err := signedHeaders(body, ts, rn)
	if err != nil {
		return nil, 0, err
	}
	return utils.Do(ctx, "POST", config.Cfg.BaseURL+path, body, h)
}

// 1.1
func GetPickupTime(ctx context.Context, serviceType uint32) (*models.GetPickupTimeResp, error) {
	req := models.GetPickupTimeReq{
		UserID:      userIDu64(),
		UserSecret:  config.Cfg.UserSecret,
		ServiceType: serviceType,
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/get_pickup_time", req)
	if err != nil {
		return nil, err
	}
	var resp models.GetPickupTimeResp
	return &resp, json.Unmarshal(b, &resp)
}

// 1.2
func CreateOrder(ctx context.Context, payload models.CreateOrderReq) (*models.CreateOrderResp, error) {
	if payload.UserID == 0 {
		payload.UserID = userIDu64()
	}
	if payload.UserSecret == "" {
		payload.UserSecret = config.Cfg.UserSecret
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_create_order", payload)
	if err != nil {
		return nil, err
	}
	var resp models.CreateOrderResp
	return &resp, json.Unmarshal(b, &resp)
}

// 1.3
func TrackOrder(ctx context.Context, payload models.TrackOrderReq) (*models.TrackOrderResp, error) {
	if payload.UserID == 0 {
		payload.UserID = userIDu64()
	}
	if payload.UserSecret == "" {
		payload.UserSecret = config.Cfg.UserSecret
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_search_order", payload)
	if err != nil {
		return nil, err
	}
	var resp models.TrackOrderResp
	return &resp, json.Unmarshal(b, &resp)
}

// 1.4
func CancelOrder(ctx context.Context, trackingNos []string) (*models.CancelOrderResp, error) {
	req := models.CancelOrderReq{
		UserID: userIDu64(), UserSecret: config.Cfg.UserSecret, TrackingNos: trackingNos,
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_cancel_order", req)
	if err != nil {
		return nil, err
	}
	var resp models.CancelOrderResp
	return &resp, json.Unmarshal(b, &resp)
}

// 1.5
func GetAWB(ctx context.Context, trackingNos []string) (*models.GetAWBResp, error) {
	req := models.GetAWBReq{
		UserID: userIDu64(), UserSecret: config.Cfg.UserSecret, TrackingNos: trackingNos,
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_get_shipping_label", req)
	if err != nil {
		return nil, err
	}
	var resp models.GetAWBResp
	return &resp, json.Unmarshal(b, &resp)
}

// 2.1
func CheckFee(ctx context.Context, payload models.CheckFeeReq) (*models.CheckFeeResp, error) {
	if payload.UserID == 0 {
		payload.UserID = userIDu64()
	}
	if payload.UserSecret == "" {
		payload.UserSecret = config.Cfg.UserSecret
	}

	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_check_order", payload)
	if err != nil {
		return nil, err
	}

	// Parse ke struct custom, bukan generic BaseResp[T]
	var resp models.CheckFeeResp
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, fmt.Errorf("SPX response parse error: %w | raw: %s", err, string(b))
	}

	// Kalau ret_code != 0, kita return error tapi tetap passing resp supaya caller bisa akses message
	if resp.RetCode != 0 {
		return &resp, fmt.Errorf("SPX error: %s", resp.Message)
	}

	return &resp, nil
}

// 2.2
func GetASF(ctx context.Context, trackingNos []string) (*models.GetASFResp, error) {
	req := models.GetASFReq{
		UserID: userIDu64(), UserSecret: config.Cfg.UserSecret, TrackingNos: trackingNos,
	}
	b, _, err := doPOST(ctx, "/open/api/v1/order/batch_get_asf", req)
	if err != nil {
		return nil, err
	}
	var resp models.GetASFResp
	return &resp, json.Unmarshal(b, &resp)
}

// 5.1
func GetAddressLink(ctx context.Context) (*models.AddressLinkResp, error) {
	emptyBody := map[string]any{} // only header required
	b, _, err := doPOST(ctx, "/open/api/v1/address/get_address_download_url", emptyBody)
	if err != nil {
		return nil, err
	}
	var resp models.AddressLinkResp
	return &resp, json.Unmarshal(b, &resp)
}
