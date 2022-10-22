package fn

// rpcf_edit_radius_account
func x10151(c conn, p Dict) Dict {
	c.Call(0x10151)
	putI(c, p, "id")
	putS(c, p, "name")
	putS(c, p, "password")
	putI(c, p, "add_nas_attrs")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
