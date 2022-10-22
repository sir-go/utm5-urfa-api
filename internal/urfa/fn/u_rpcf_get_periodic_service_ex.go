package fn

// rpcf_get_periodic_service_ex
func x2230(c conn, p Dict) Dict {
	c.Call(0x2230)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"service_name":    c.GetS(),
		"comment":         c.GetS(),
		"link_by_default": c.GetI(),
		"contract_type":   c.GetI(),
		"cost":            c.GetD(),
		"deprecated":      c.GetI(),
		"discount_method": c.GetI(),
		"scheme_id":       c.GetI(),
		"param":           c.GetI(),
		"tariff_id":       c.GetI(),
		"parent_id":       c.GetI(),
	}
}
