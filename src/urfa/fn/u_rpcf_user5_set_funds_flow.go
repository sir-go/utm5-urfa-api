package fn

// rpcf_user5_set_funds_flow
func xU15007(c conn, p Dict) Dict {
	c.Call(-0x15007)
	putI(c, p, "account_id_from")
	putI(c, p, "account_id_to")
	putD(c, p, "amount")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
