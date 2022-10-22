package fn

// rpcf_edit_payment_method
func x3102(c conn, p Dict) Dict {
	c.Call(0x3102)
	putI(c, p, "id")
	putS(c, p, "name")
	c.Send()

	return nil
}
