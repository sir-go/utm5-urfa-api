package fn

// rpcf_add_payment_method
func x3101(c conn, p Dict) Dict {
	c.Call(0x3101)
	putI(c, p, "id")
	putS(c, p, "name")
	c.Send()

	return nil
}
