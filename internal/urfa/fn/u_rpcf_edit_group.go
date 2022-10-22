package fn

// rpcf_edit_group
func x2402(c conn, p Dict) Dict {
	c.Call(0x2402)
	putI(c, p, "group_id")
	putS(c, p, "group_name")
	c.Send()

	return nil
}
