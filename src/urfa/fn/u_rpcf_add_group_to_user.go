package fn

// rpcf_add_group_to_user
func x2552(c conn, p Dict) Dict {
	c.Call(0x2552)
	putI(c, p, "user_id")
	putI(c, p, "group_id")
	c.Send()

	return nil
}
