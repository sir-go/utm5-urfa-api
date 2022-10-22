package fn

// rpcf_remove_street
func x2819(c conn, p Dict) Dict {
	c.Call(0x2819)
	putI(c, p, "street_id")
	c.Send()

	return nil
}
