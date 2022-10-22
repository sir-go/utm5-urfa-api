package fn

// rpcf_change_account_balance
func x2045(c conn, p Dict) Dict {
	c.Call(0x2045)
	putI(c, p, "account_id")
	putD(c, p, "balance")
	putS(c, p, "comment")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
