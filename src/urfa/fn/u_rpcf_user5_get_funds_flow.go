package fn

// rpcf_user5_get_funds_flow
func xU15008(c conn, p Dict) Dict {
	c.Call(-0x15008)
	putI(c, p, "unused")
	putI(c, p, "uid")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
