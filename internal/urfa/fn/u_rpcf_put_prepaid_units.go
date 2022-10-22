package fn

// rpcf_put_prepaid_units
func x5501(c conn, p Dict) Dict {
	c.Call(0x5501)
	putI(c, p, "slink_id")
	putI(c, p, "tclass_id")
	putL(c, p, "prepaid_units")
	c.Send()

	return nil
}
