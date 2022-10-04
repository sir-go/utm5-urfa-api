package fn

// rpcf_user5_set_promised_payment
func xU15025(c conn, p Dict) Dict {
	c.Call(-0x15025)
	putI(c, p, "account_id", 0)
	putD(c, p, "value")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
