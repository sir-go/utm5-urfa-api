package fn

// rpcf_user5_get_requirement_payment
func xU15060(c conn, p Dict) Dict {
	c.Call(-0x15060)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"result": c.GetD(),
	}
}
