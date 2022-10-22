package fn

// rpcf_edit_funds_flow_settings
func x15026(c conn, p Dict) Dict {
	c.Call(0x15026)
	putI(c, p, "id")
	putI(c, p, "group_id")
	putI(c, p, "priopity")
	putI(c, p, "is_enabled")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
