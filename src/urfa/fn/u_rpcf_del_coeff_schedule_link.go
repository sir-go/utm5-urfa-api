package fn

// rpcf_del_coeff_schedule_link
func x4714(c conn, p Dict) Dict {
	c.Call(0x4714)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
