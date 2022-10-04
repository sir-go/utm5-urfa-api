package fn

// rpcf_add_group
func x2401(c conn, p Dict) Dict {
	c.Call(0x2401)
	putI(c, p, "group_id", 0)
	putS(c, p, "group_name")
	c.Send()

	return nil
}
