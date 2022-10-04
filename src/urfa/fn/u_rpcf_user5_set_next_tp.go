package fn

// rpcf_user5_set_next_tp
func xU15006(c conn, p Dict) Dict {
	c.Call(-0x15006)
	putI(c, p, "account_id", 0)
	putI(c, p, "tp_link", 0)
	putI(c, p, "tp_next", 0)
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
