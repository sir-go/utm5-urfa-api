package fn

// rpcf_edit_currency
func x2912(c conn, p Dict) Dict {
	c.Call(0x2912)
	putI(c, p, "id")
	putS(c, p, "currency_brief_name")
	putS(c, p, "currency_full_name")
	putI(c, p, "date")
	putD(c, p, "rate")
	putD(c, p, "percent")
	c.Send()

	return nil
}
