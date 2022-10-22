package fn

// rpcf_del_group
func x240b(c conn, p Dict) Dict {
	c.Call(0x240b)
	putI(c, p, "group_id")
	c.Send()

	return nil
}
