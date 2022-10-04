package fn

// rpcf_get_vod_service_link
func x4516(c conn, p Dict) Dict {
	c.Call(0x4516)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"slink_id":    c.GetI(),
		"ap_id":       c.GetI(),
		"start_date":  c.GetI(),
		"expire_date": c.GetI(),
		"cost_coef":   c.GetD(),
	}
}
