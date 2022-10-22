package fn

// rpcf_edit_tariff
func x3013(c conn, p Dict) Dict {
	c.Call(0x3013)
	putI(c, p, "tp_id")
	putS(c, p, "name")
	putI(c, p, "expire_date")
	putI(c, p, "is_blocked")
	putI(c, p, "balance_rollover")
	c.Send()

	return Dict{
		"tp_id": c.GetI(),
	}
}
