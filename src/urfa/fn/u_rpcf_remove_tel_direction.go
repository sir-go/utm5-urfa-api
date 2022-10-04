package fn

// rpcf_remove_tel_direction
func x10302(c conn, p Dict) Dict {
	c.Call(0x10302)
	putI(c, p, "dir_id")
	c.Send()

	return Dict{
		"dir_id": c.GetI(),
	}
}
