package models

type (
	Order struct {
		order_uid          string
		track_bumber       string
		entry              string
		delivery           Delivery
		payment            Payment
		items              []Item
		locale             string
		internal_signature string
		customer_id        string
		delivery_service   string
		shardkey           string
		sm_id              int
		date_created       string
		oof_shard          string
	}

	Delivery struct {
		name    string
		phone   string
		zip     string
		city    string
		address string
		region  string
		email   string
	}

	Payment struct {
		transaction   string
		request_id    string
		currency      string
		provider      string
		amount        int
		payment_dt    int
		bank          string
		delivery_cost int
		goods_total   int
		custom_fee    int
	}

	Item struct {
		chrt_id      int
		track_number string
		price        int
		rid          string
		name         string
		sale         int
		size         string
		total_price  int
		nm_id        int
		brand        string
		status       int16
	}
)
