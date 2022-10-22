package fn

// rpcf_edit_vod_service_link
func x4515(c conn, p Dict) Dict {
	c.Call(0x4515)
	putI(c, p, "slink_id")
	putI(c, p, "start_date")
	putD(c, p, "cost_coef", 1.0)
	c.Send()

	return Dict{
		"slink_id": c.GetI(),
	}
}
