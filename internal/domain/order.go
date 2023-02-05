package domain

type Order struct {
	OrderUID    string `json:"order_uid," validate:"required,min=19,max=19"`
	TrackNumber string `json:"track_number" validate:"required,min=14,max=14"`
	Entry       string `json:"entry" validate:"required,min=4,max=4 "`
	Delivery    struct {
		Name    string `json:"name" validate:"required,max=40"`
		Phone   string `json:"phone" validate:"required,max=11"`
		Zip     string `json:"zip" validate:"required,max=7"`
		City    string `json:"city" validate:"required,max=20"`
		Address string `json:"address" validate:"required,max=30"`
		Region  string `json:"region" validate:"required,max=20"`
		Email   string `json:"email" validate:"required,max=50"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction" validate:"required"`
		RequestId    string `json:"request_id" validate:"required"`
		Currency     string `json:"currency" validate:"required"`
		Provider     string `json:"provider" validate:"required"`
		Amount       int    `json:"amount" validate:"required"`
		PaymentDt    int    `json:"payment_dt" validate:"required"`
		Bank         string `json:"bank" validate:"required"`
		DeliveryCost int    `json:"delivery_cost" validate:"required"`
		GoodsTotal   int    `json:"goods_total" validate:"required"`
		CustomFee    int    `json:"custom_fee" validate:"required"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id" validate:"required"`
		TrackNumber string `json:"track_number" validate:"required,min=14,max=14"`
		Price       int    `json:"price" validate:"required"`
		Rid         string `json:"rid" validate:"required,max=21"`
		Name        string `json:"name" validate:"required"`
		Sale        int    `json:"sale" validate:"required"`
		Size        string `json:"size" validate:"required"`
		TotalPrice  int    `json:"total_price" validate:"required"`
		NmId        int    `json:"nm_id" validate:"required"`
		Brand       string `json:"brand" validate:"required"`
		Status      int    `json:"status" validate:"required,max=999"`
	} `json:"items"`
	Locale            string `json:"locale" validate:"oneof=ru en it"`
	InternalSignature string `json:"internal_signature" validate:"required,min=8,max=8"`
	CustomerId        string `json:"customer_id" validate:"required,min=4,max=4"`
	DeliveryService   string `json:"delivery_service" validate:"required,min=5,max=5"`
	Shardkey          string `json:"shardkey" validate:"required,max=2"`
	SmId              int    `json:"sm_id" validate:"required,gte=0,lte=5000"`
	DateCreated       string `json:"date_created" validate:"required"`
	OofShard          string `json:"oof_shard" validate:"required,max=2"`
}
