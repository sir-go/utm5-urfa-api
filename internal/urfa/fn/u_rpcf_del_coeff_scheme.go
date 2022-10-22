package fn

// rpcf_del_coeff_scheme
func x4704(c conn, p Dict) Dict {
	c.Call(0x4704)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
