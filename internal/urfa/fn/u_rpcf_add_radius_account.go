package fn

// rpcf_add_radius_account
func x10152(c conn, p Dict) Dict {
	c.Call(0x10152)
	putS(c, p, "name")
	putS(c, p, "password")
	putI(c, p, "add_nas_attrs")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
