package fn

// rpcf_add_iptv_service_link
func x4504(c conn, p Dict) Dict {
	c.Call(0x4504)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "service_id")
	putI(c, p, "tp_link_id")
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
