package fn

// rpcf_remove_service
func x210e(c conn, p Dict) Dict {
	c.Call(0x210e)
	putI(c, p, "sid")
	c.Send()

	return nil
}
