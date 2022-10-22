package fn

// rpcf_get_group_id_by_name
func x240c(c conn, p Dict) Dict {
	c.Call(0x240c)
	putS(c, p, "group_name")
	c.Send()

	return Dict{
		"group_id": c.GetI(),
	}
}
