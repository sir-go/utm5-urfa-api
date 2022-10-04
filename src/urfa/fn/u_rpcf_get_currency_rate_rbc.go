package fn

// rpcf_get_currency_rate_rbc
func x2914(c conn, p Dict) Dict {
	c.Call(0x2914)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"rate":  c.GetD(),
		"error": c.GetS(),
	}
}
