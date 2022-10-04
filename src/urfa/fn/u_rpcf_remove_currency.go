package fn

// rpcf_remove_currency
func x2915(c conn, p Dict) Dict {
	c.Call(0x2915)
	putI(c, p, "id")
	c.Send()

	return nil
}
