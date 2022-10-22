package fn

// rpcf_del_coeff_schedule
func x4709(c conn, p Dict) Dict {
	c.Call(0x4709)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
