package fn

// rpcf_remove_switch_type
func x502(c conn, p Dict) Dict {
	c.Call(0x502)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
