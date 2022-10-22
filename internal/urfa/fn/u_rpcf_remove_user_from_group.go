package fn

// rpcf_remove_user_from_group
func x2408(c conn, p Dict) Dict {
	c.Call(0x2408)
	putI(c, p, "user_id")
	putI(c, p, "group_id")
	c.Send()

	return nil
}
