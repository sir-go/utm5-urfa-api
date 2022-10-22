package fn

// rpcf_get_once_service_new
func x2115(c conn, p Dict) Dict {
	c.Call(0x2115)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"service_name":    c.GetS(),
		"comment":         c.GetS(),
		"link_by_default": c.GetI(),
		"cost":            c.GetD(),
		"drop_from_group": c.GetI(),
		"is_parent_id":    c.GetI(),
		"tariff_id":       c.GetI(),
		"parent_id":       c.GetI(),
	}
}
