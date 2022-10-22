package fn

// rpcf_request_traffic_detailed
func x502a(c conn, p Dict) Dict {
	c.Call(0x502a)
	putI(c, p, "user_id", 0)
	putI(c, p, "account_id")
	putI(c, p, "collector_id")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	return Dict{
		"is_immediate": c.GetI(),
	}
}
