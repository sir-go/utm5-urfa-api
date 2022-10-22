package fn

// rpcf_remove_dhcp_options
func x1102(c conn, p Dict) Dict {
	c.Call(0x1102)
	putI(c, p, "id")
	putI(c, p, "type")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
