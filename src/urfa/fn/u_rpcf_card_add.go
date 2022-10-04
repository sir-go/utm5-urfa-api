package fn

// rpcf_card_add
func x4202(c conn, p Dict) Dict {
	c.Call(0x4202)
	putS(c, p, "secret")
	putD(c, p, "balance")
	putI(c, p, "currency")
	putI(c, p, "expire")
	putI(c, p, "days")
	putI(c, p, "is_used")
	putI(c, p, "tp_id")
	c.Send()

	return nil
}
