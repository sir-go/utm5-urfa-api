package fn

// rpcf_remove_house
func x2815(c conn, p Dict) Dict {
	c.Call(0x2815)
	putI(c, p, "hid")
	c.Send()

	return nil
}
