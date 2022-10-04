package fn

// rpcf_user5_get_group_id_by_name
func xU401b(c conn, p Dict) Dict {
	c.Call(-0x401b)
	putS(c, p, "name")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
