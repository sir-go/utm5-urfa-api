package fn

// rpcf_make_lease_overdue
func x1351(c conn, p Dict) Dict {
	c.Call(0x1351)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"ret_code": c.GetI(),
	}
}
