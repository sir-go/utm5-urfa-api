package fn

// rpcf_get_users_count
func x2011(c conn, p Dict) Dict {
	c.Call(0x2011)
	putI(c, p, "card_user", 0)
	c.Send()

	return Dict{
		"count": c.GetI(),
	}
}
