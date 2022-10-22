package fn

// rpcf_edit_ippool
func x1069(c conn, p Dict) Dict {
	c.Call(0x1069)
	putI(c, p, "id")
	putS(c, p, "name")
	putA(c, p, "address")
	putI(c, p, "mask")
	c.Send()

	return nil
}
