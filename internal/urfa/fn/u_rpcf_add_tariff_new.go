package fn

// rpcf_add_tariff_new
func x3041(c conn, p Dict) Dict {
	c.Call(0x3041)
	putS(c, p, "name")
	putI(c, p, "balance_rollover")
	putS(c, p, "comments")
	c.Send()

	return Dict{
		"tp_id": c.GetI(),
	}
}
