package fn

// rpcf_get_sysuser_name
func x4408(c conn, p Dict) Dict {
	c.Call(0x4408)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"login": c.GetS(),
	}
}
