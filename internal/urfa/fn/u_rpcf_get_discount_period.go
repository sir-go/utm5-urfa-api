package fn

// rpcf_get_discount_period
func x2609(c conn, p Dict) Dict {
	c.Call(0x2609)
	putI(c, p, "discount_period_id")
	c.Send()

	return Dict{
		"start_date":              c.GetI(),
		"end_date":                c.GetI(),
		"periodic_type":           c.GetI(),
		"custom_duration":         c.GetI(),
		"discounts_per_week":      c.GetI(),
		"next_discount_period_id": c.GetI(),
		"invoice_month":           c.GetI(),
	}
}
