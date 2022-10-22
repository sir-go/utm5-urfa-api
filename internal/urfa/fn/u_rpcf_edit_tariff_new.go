package fn

// rpcf_edit_tariff_new
func x3042(c conn, p Dict) Dict {
	c.Call(0x3042)
	putI(c, p, "tp_id")
	putS(c, p, "name")
	putI(c, p, "balance_rollover")
	putS(c, p, "comments")
	c.Send()

	return Dict{
		"tp_id": c.GetI(),
	}
}
