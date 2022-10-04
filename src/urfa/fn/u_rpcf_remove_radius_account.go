package fn

// rpcf_remove_radius_account
func x10153(c conn, p Dict) Dict {
	c.Call(0x10153)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
