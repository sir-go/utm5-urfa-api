package fn

// rpcf_get_periodic_service
func x2244(c conn, p Dict) Dict {
	c.Call(0x2244)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"service_name":    c.GetS(),
		"comment":         c.GetS(),
		"link_by_default": c.GetI(),
		"cost":            c.GetD(),
		"deprecated":      c.GetI(),
		"discount_method": c.GetI(),
		"scheme_id":       c.GetI(),
		"start_date":      c.GetI(),
		"expire_date":     c.GetI(),
		"param":           c.GetI(),
		"tariff_id":       c.GetI(),
		"parent_id":       c.GetI(),
		"contract_type":   c.GetI(),
	}
}
