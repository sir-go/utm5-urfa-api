package fn

// rpcf_add_tariff
func x3012(c conn, p Dict) Dict {
	c.Call(0x3012)
	putS(c, p, "name")
	putI(c, p, "expire_date")
	putI(c, p, "is_blocked")
	putI(c, p, "balance_rollover")
	c.Send()

	return Dict{
		"tp_id": c.GetI(),
	}
}
