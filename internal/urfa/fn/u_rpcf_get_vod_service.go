package fn

// rpcf_get_vod_service
func x4532(c conn, p Dict) Dict {
	c.Call(0x4532)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"service_id":      c.GetI(),
		"parent_id":       c.GetI(),
		"tariff_id":       c.GetI(),
		"service_name":    c.GetS(),
		"comment":         c.GetS(),
		"link_by_default": c.GetI(),
		"is_dynamic":      c.GetI(),
		"contract_type":   c.GetI(),
		"cost":            c.GetD(),
		"duration":        c.GetI(),
		"tv_type":         c.GetI(),
		"media_content":   c.GetI(),
		"media_group":     c.GetL(),
		"custom_data":     c.GetS(),
	}
}
