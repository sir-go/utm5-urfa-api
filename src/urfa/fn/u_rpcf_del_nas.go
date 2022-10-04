package fn

// rpcf_del_nas
func x501a(c conn, p Dict) Dict {
	c.Call(0x501a)
	putI(c, p, "id")
	c.Send()

	return nil
}
