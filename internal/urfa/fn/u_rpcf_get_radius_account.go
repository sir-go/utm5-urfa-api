package fn

// rpcf_get_radius_account
func x10150(c conn, p Dict) Dict {
	c.Call(0x10150)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":            c.GetI(),
		"name":          c.GetS(),
		"password":      c.GetS(),
		"add_nas_attrs": c.GetI(),
	}
}
