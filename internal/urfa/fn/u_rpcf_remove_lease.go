package fn

// rpcf_remove_lease
func x1352(c conn, p Dict) Dict {
	c.Call(0x1352)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"ret_code": c.GetI(),
	}
}
