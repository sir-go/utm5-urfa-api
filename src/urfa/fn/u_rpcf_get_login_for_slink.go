package fn

// rpcf_get_login_for_slink
func x200d(c conn, p Dict) Dict {
	c.Call(0x200d)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"user_data_full_login": c.GetS(),
	}
}
