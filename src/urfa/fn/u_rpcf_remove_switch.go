package fn

// rpcf_remove_switch
func x1152(c conn, p Dict) Dict {
	c.Call(0x1152)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
