package models

// 2.1 Check Shipping Fee
type CheckFeeReq struct {
	UserID     uint64        `json:"user_id"`
	UserSecret string        `json:"user_secret"`
	Orders     []CreateOrder `json:"orders"`
}

type CheckFeeOrder struct {
	EstimatedShippingFee            FloatString `json:"estimated_shipping_fee"`
	BasicShippingFee                FloatString `json:"basic_shipping_fee"`
	CODServiceFee                   FloatString `json:"cod_service_fee"`
	InsuranceServiceFee             FloatString `json:"insurance_service_fee"`
	VATFee                          FloatString `json:"vat_fee"`
	EstimatedDistance               FloatString `json:"estimated_distance"`
	BasicShippingFeeWithoutDiscount FloatString `json:"basic_shipping_fee_without_discount"`
}

type CheckFeeResp struct {
	RetCode int               `json:"ret_code"`
	Message string            `json:"message"`
	Data    *CheckFeeRespData `json:"data,omitempty"`
}

type CheckFeeRespData struct {
	Orders   []CheckFeeOrder `json:"orders"`
	FailList []struct {
		RetCode  int    `json:"ret_code"`
		Message  string `json:"message"`
		DebugMsg string `json:"debug_msg"`
		OrderID  string `json:"order_id"`
	} `json:"fail_list"`
}

// 2.2 Get Order Fee (ASF)
type GetASFReq struct {
	UserID      uint64   `json:"user_id"`
	UserSecret  string   `json:"user_secret"`
	TrackingNos []string `json:"tracking_no_list"`
}
type ASFOrder struct {
	TrackingNo string `json:"tracking_no"`
	ParcelInfo struct {
		ActualWeight        float64 `json:"actual_weight"`
		ActualLength        float64 `json:"actual_length"`
		ActualWidth         float64 `json:"actual_width"`
		ActualHeight        float64 `json:"actual_height"`
		ChargeableWeight    float64 `json:"chargeable_weight"`
		InsuranceCollection int     `json:"insurance_collection"`
		ExpressInsuredValue float64 `json:"express_insured_value"`
		CODCollection       int     `json:"cod_collection"`
		CODAmount           float64 `json:"cod_amount"`
	} `json:"parcel_info"`
	FeeInfo struct {
		PaymentStatus       int     `json:"payment_status"`
		PaymentRole         int     `json:"payment_role"`
		ActualShippingFee   float64 `json:"actual_shipping_fee"`
		ReturnShippingFee   float64 `json:"return_shipping_fee"`
		BasicShippingFee    float64 `json:"basic_shipping_fee"`
		InsuranceServiceFee float64 `json:"insurance_service_fee"`
		CODServiceFee       float64 `json:"cod_service_fee"`
	} `json:"fee_info"`
	DeliverInfo struct {
		ActualDistance float64 `json:"actual_distance"`
	} `json:"deliver_info"`
}
type GetASFRespData struct {
	Orders   []ASFOrder `json:"orders"`
	FailList []struct {
		RetCode    int    `json:"ret_code"`
		Message    string `json:"message"`
		TrackingNo string `json:"tracking_no"`
	} `json:"fail_list"`
}
type GetASFResp = BaseResp[GetASFRespData]
