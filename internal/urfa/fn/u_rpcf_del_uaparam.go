package fn

// rpcf_del_uaparam
func x440d(c conn, p Dict) Dict {
	c.Call(0x440d)
	putI(c, p, "uaparam_id")
	c.Send()

	return nil
}
