package models

// 1.1 Get Pickup Timeslot (req/resp)
type GetPickupTimeReq struct {
	UserID      uint64 `json:"user_id"`
	UserSecret  string `json:"user_secret"`
	ServiceType uint32 `json:"service_type"`
}
type PickupSlot struct {
	Date       string      `json:"date"`
	PickupTime uint64      `json:"pickup_time"`
	Slots      interface{} `json:"slots"` // bisa spesifikin kalau mau
}
type GetPickupTimeResp = BaseResp[[]PickupSlot]

// 1.2 Create Order
type CreateOrderReq struct {
	UserID     uint64        `json:"user_id"`
	UserSecret string        `json:"user_secret"`
	Orders     []CreateOrder `json:"orders"`
}
type CreateOrder struct {
	OrderID  string      `json:"order_id,omitempty"`
	BaseInfo BaseInfo    `json:"base_info"`
	Sender   SenderInfo  `json:"sender_info"`
	Fullfill Fulfillment `json:"fulfillment_info"`
	Deliver  DeliverInfo `json:"deliver_info"`
	Parcel   ParcelInfo  `json:"parcel_info"`
}
type BaseInfo struct {
	ServiceType int `json:"service_type"`
}
type SenderInfo struct {
	SenderState         string `json:"sender_state"`
	SenderCity          string `json:"sender_city"`
	SenderDistrict      string `json:"sender_district"`
	SenderLongitude     string `json:"sender_longitude,omitempty"`
	SenderLatitude      string `json:"sender_latitude,omitempty"`
	SenderName          string `json:"sender_name"`
	SenderPhone         string `json:"sender_phone"`
	SenderDetailAddress string `json:"sender_detail_address"`
}
type Fulfillment struct {
	PaymentRole         int     `json:"payment_role"`
	CODCollection       int     `json:"cod_collection"`
	CODAmount           float64 `json:"cod_amount,omitempty"`
	InsuranceCollection int     `json:"insurance_collection"`
	CollectType         int     `json:"collect_type"`
	PickupTime          uint64  `json:"pickup_time,omitempty"`
	PickupTimeRangeID   uint64  `json:"pickup_time_range_id,omitempty"`
	PickupTimeRange     string  `json:"pickup_time_range,omitempty"`
	AllowMutualCheck    int     `json:"allow_mutual_check,omitempty"`
}
type DeliverInfo struct {
	DeliverState         string `json:"deliver_state"`
	DeliverCity          string `json:"deliver_city"`
	DeliverDistrict      string `json:"deliver_district"`
	DeliverLongitude     string `json:"deliver_longitude,omitempty"`
	DeliverLatitude      string `json:"deliver_latitude,omitempty"`
	DeliverName          string `json:"deliver_name"`
	DeliverPhone         string `json:"deliver_phone"`
	DeliverDetailAddress string `json:"deliver_detail_address"`
	DeliverInstruction   string `json:"deliver_instruction,omitempty"`
}
type ParcelInfo struct {
	ParcelWeight        float64 `json:"parcel_weight"`
	ParcelLength        float64 `json:"parcel_length,omitempty"`
	ParcelWidth         float64 `json:"parcel_width,omitempty"`
	ParcelHeight        float64 `json:"parcel_height,omitempty"`
	ParcelItemName      string  `json:"parcel_item_name"`
	ParcelItemQuantity  int     `json:"parcel_item_quantity"`
	ExpressInsuredValue float64 `json:"express_insured_value"`
}

type CreateOrderRespData struct {
	Orders   []CreatedOrder `json:"orders"`
	FailList []CreateFail   `json:"fail_list"`
}
type CreatedOrder struct {
	OrderID              string  `json:"order_id"`
	OrderIDLink          string  `json:"order_id_link"`
	TrackingNo           string  `json:"tracking_no"`
	TrackingLink         string  `json:"tracking_link"`
	RFirstSortCode       string  `json:"r_first_sort_code"`
	RThirdSortCode       string  `json:"r_third_sort_code"`
	ReturnFirstSortCode  string  `json:"r_return_sort_code"`
	EstimatedShippingFee float64 `json:"estimated_shipping_fee"`
	BasicShippingFee     float64 `json:"basic_shipping_fee"`
	CODServiceFee        float64 `json:"cod_service_fee"`
	InsuranceServiceFee  float64 `json:"insurance_service_fee"`
	VATFee               float64 `json:"vat_fee"`
}
type CreateFail struct {
	RetCode  int    `json:"ret_code"`
	Message  string `json:"message"`
	DebugMsg string `json:"debug_msg"`
	OrderID  string `json:"order_id"`
}
type CreateOrderResp = BaseResp[CreateOrderRespData]

// 1.3 Track Order
type TrackOrderReq struct {
	UserID      uint64   `json:"user_id"`
	UserSecret  string   `json:"user_secret"`
	TrackingNos []string `json:"tracking_no_list,omitempty"`
	OrderIDs    []string `json:"order_id_list,omitempty"`
	BatchNo     uint64   `json:"batch_no,omitempty"`
}
type Route struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Timestamp  uint64 `json:"timestamp"`
}
type TrackedOrder struct {
	TrackingNo   string      `json:"tracking_no"`
	TrackingLink string      `json:"tracking_link"`
	OrderID      string      `json:"order_id"`
	OrderIDLink  string      `json:"order_id_link"`
	Status       string      `json:"status"`
	StatusCode   string      `json:"status_code"`
	BaseInfo     BaseInfo    `json:"base_info"`
	Sender       interface{} `json:"sender_info"`
	Fulfillment  interface{} `json:"fulfillment_info"`
	Deliver      interface{} `json:"deliver_info"`
	Parcel       interface{} `json:"parcel_info"`
	Routes       []Route     `json:"routes"`
}
type TrackOrderRespData struct {
	Orders   []TrackedOrder `json:"orders"`
	FailList []struct {
		RetCode    int    `json:"ret_code"`
		Message    string `json:"message"`
		TrackingNo string `json:"tracking_no"`
		OrderID    string `json:"order_id"`
	} `json:"fail_list"`
}
type TrackOrderResp = BaseResp[TrackOrderRespData]

// 1.4 Cancel Order
type CancelOrderReq struct {
	UserID      uint64   `json:"user_id"`
	UserSecret  string   `json:"user_secret"`
	TrackingNos []string `json:"tracking_no_list"`
}
type CancelOrderRespData struct {
	TrackingNoList []string `json:"tracking_no_list"`
	FailList       []struct {
		RetCode    int    `json:"ret_code"`
		Message    string `json:"message"`
		TrackingNo string `json:"tracking_no"`
	} `json:"fail_list"`
}
type CancelOrderResp = BaseResp[CancelOrderRespData]

// 1.5 Get AWB
type GetAWBReq struct {
	UserID      uint64   `json:"user_id"`
	UserSecret  string   `json:"user_secret"`
	TrackingNos []string `json:"tracking_no_list"`
}
type GetAWBRespData struct {
	AWBLink  string `json:"awb_link"`
	FailList []struct {
		RetCode    int    `json:"ret_code"`
		Message    string `json:"message"`
		TrackingNo string `json:"tracking_no"`
	} `json:"fail_list"`
}
type GetAWBResp = BaseResp[GetAWBRespData]

// 5.1 Address file link
type AddressLinkRespData struct {
	AddressDownloadURL string `json:"address_download_url"`
}
type AddressLinkResp = BaseResp[AddressLinkRespData]
