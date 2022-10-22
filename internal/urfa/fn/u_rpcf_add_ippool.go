package fn

// rpcf_add_ippool
func x1068(c conn, p Dict) Dict {
	c.Call(0x1068)
	putI(c, p, "id")
	putS(c, p, "name")
	putI(c, p, "address")
	putI(c, p, "mask")
	c.Send()

	return nil
}
