package fn

// rpcf_dealer_change_password_self
func x7000004f(c conn, p Dict) Dict {
	c.Call(0x7000004f)
	putS(c, p, "old_pass")
	putS(c, p, "new_pass")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
