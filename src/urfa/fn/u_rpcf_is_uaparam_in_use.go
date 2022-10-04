package fn

// rpcf_is_uaparam_in_use
func x4412(c conn, p Dict) Dict {
	c.Call(0x4412)
	putI(c, p, "uaparam_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
