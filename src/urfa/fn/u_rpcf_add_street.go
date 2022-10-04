package fn

// rpcf_add_street
func x2817(c conn, p Dict) Dict {
	c.Call(0x2817)
	putI(c, p, "street_id", 0)
	putS(c, p, "country", "РФ")
	putS(c, p, "region", "")
	putS(c, p, "city", "")
	putS(c, p, "street")
	putS(c, p, "number", "")
	c.Send()

	return nil
}
