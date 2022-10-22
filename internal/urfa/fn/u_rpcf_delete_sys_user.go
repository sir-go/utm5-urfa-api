package fn

// rpcf_delete_sys_user
func x4410(c conn, p Dict) Dict {
	c.Call(0x4410)
	putI(c, p, "user_id")
	c.Send()

	return nil
}
