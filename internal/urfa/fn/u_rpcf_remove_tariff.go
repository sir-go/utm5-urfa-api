package fn

// rpcf_remove_tariff
func x301b(c conn, p Dict) Dict {
	c.Call(0x301b)
	putI(c, p, "tid")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
