package fn

// rpcf_get_currency_rate_to_date
func x8003(c conn, p Dict) Dict {
	c.Call(0x8003)
	putI(c, p, "date")
	putI(c, p, "code")
	c.Send()

	return Dict{
		"result": c.GetD(),
	}
}
