package fn

// rpcf_add_funds_flow_settings
func x15025(c conn, p Dict) Dict {
	c.Call(0x15025)
	putI(c, p, "group_id")
	putI(c, p, "priopity")
	putI(c, p, "is_enabled")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
