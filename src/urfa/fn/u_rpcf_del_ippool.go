package fn

// rpcf_del_ippool
func x1064(c conn, p Dict) Dict {
	c.Call(0x1064)
	putI(c, p, "id")
	putS(c, p, "name")
	putA(c, p, "address")
	putI(c, p, "mask")
	c.Send()

	return nil
}
