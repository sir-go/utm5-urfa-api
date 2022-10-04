package fn

// rpcf_get_hotspot_service_link
func x272c(c conn, p Dict) Dict {
	c.Call(0x272c)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"tariff_link_id":     c.GetI(),
		"is_blocked":         c.GetI(),
		"discount_period_id": c.GetI(),
		"start_date":         c.GetI(),
		"expire_date":        c.GetI(),
		"policy_id":          c.GetI(),
		"cost_coef":          c.GetD(),
		"house_id":           c.GetI(),
		"comment":            c.GetS(),
		"login":              c.GetS(),
		"password":           c.GetS(),
		"is_unabon_period":   c.GetI(),
		"is_unprepay_period": c.GetI(),
		"tariff_id":          c.GetI(),
		"parent_id":          c.GetI(),
	}
}
