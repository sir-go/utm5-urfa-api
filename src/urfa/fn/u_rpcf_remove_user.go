package fn

// rpcf_remove_user
func x200e(c conn, p Dict) Dict {
	c.Call(0x200e)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
