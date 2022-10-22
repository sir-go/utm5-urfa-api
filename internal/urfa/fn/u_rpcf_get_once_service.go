package fn

// rpcf_get_once_service
func x210a(c conn, p Dict) Dict {
	c.Call(0x210a)
	putI(c, p, "sid")
	c.Send()

	return Dict{
		"service_name":    c.GetS(),
		"comment":         c.GetS(),
		"link_by_default": c.GetI(),
		"cost":            c.GetD(),
		"is_parent_id":    c.GetI(),
		"tariff_id":       c.GetI(),
		"parent_id":       c.GetI(),
	}
}
