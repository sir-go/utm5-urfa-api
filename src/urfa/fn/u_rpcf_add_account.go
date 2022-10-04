package fn

// rpcf_add_account
func x1510a(c conn, p Dict) Dict {
	c.Call(0x1510a)
	putI(c, p, "user_id")
	putI(c, p, "is_basic", 1)
	putI(c, p, "is_blocked", 0)
	putD(c, p, "balance", 0)
	putD(c, p, "credit", 0.0)
	putD(c, p, "vat_rate", 0.0)
	putD(c, p, "sale_tax_rate", 0.0)
	putI(c, p, "int_status", 1)
	putI(c, p, "unlimited", 0)
	putI(c, p, "auto_enable_inet", 1)
	putS(c, p, "external_id", "")
	c.Send()

	if aid := c.GetI(); aid != 0 {
		return Dict{"account_id": aid}
	}

	return Dict{"error": Dict{"code": 11, "comment": "unable to add account"}}
}
