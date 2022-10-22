package fn

// rpcf_del_uaparam_new
func x4411(c conn, p Dict) Dict {
	c.Call(0x4411)
	putI(c, p, "uaparam_id")
	c.Send()

	return nil
}
