package fn

// rpcf_get_iptv_service_link
func x4506(c conn, p Dict) Dict {
	c.Call(0x4506)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"slink_id":    c.GetI(),
		"ap_id":       c.GetI(),
		"start_date":  c.GetI(),
		"expire_date": c.GetI(),
		"policy_id":   c.GetI(),
		"cost_coef":   c.GetD(),
	}
}
