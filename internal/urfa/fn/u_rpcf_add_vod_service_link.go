package fn

// rpcf_add_vod_service_link
func x4514(c conn, p Dict) Dict {
	c.Call(0x4514)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "service_id")
	putI(c, p, "tp_link_id")
	putI(c, p, "start_date")
	putD(c, p, "cost_coef", 1.0)
	c.Send()

	return Dict{
		"slink_id": c.GetI(),
	}
}
