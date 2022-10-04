package fn

// rpcf_change_own_password
func x2124(c conn, p Dict) Dict {
	c.Call(0x2124)
	putS(c, p, "old_pass")
	putS(c, p, "new_pass")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
