package fn

// rpcf_remove_dhcp_pool
func x702(c conn, p Dict) Dict {
	c.Call(0x702)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
