package fn

// rpcf_add_supplier
func x8016(c conn, p Dict) Dict {
	c.Call(0x8016)
	putS(c, p, "name")
	putS(c, p, "short_name")
	putS(c, p, "jur_adress")
	putS(c, p, "act_adress")
	putS(c, p, "inn")
	putS(c, p, "kpp")
	putI(c, p, "bank_id")
	putS(c, p, "account")
	putS(c, p, "headman")
	putS(c, p, "bookeeper")
	putS(c, p, "short_headman")
	putS(c, p, "short_bookeeper")
	putD(c, p, "supp_balance")
	putD(c, p, "tax_rate")
	putI(c, p, "type")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
