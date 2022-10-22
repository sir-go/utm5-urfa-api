package fn

// rpcf_get_requirement_payment
func x15111(c conn, p Dict) Dict {
	c.Call(0x15111)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"result": c.GetD(),
	}
}
