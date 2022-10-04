package fn

// rpcf_edit_iptv_service_link
func x4505(c conn, p Dict) Dict {
	c.Call(0x4505)
	putI(c, p, "slink_id")
	putI(c, p, "ap_id")
	putI(c, p, "start_date")
	putI(c, p, "expire_date")
	putI(c, p, "policy_id")
	putD(c, p, "cost_coef", 1.0)
	c.Send()

	return Dict{
		"slink_id": c.GetI(),
	}
}
