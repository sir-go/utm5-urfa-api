package fn

// rpcf_delete_funds_flow_settings
func x15027(c conn, p Dict) Dict {
	c.Call(0x15027)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
