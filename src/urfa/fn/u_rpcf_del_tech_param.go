package fn

// rpcf_del_tech_param
func x9001(c conn, p Dict) Dict {
	c.Call(0x9001)
	putI(c, p, "tpid")
	putI(c, p, "slink")
	c.Send()

	return nil
}
