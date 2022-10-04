package fn

// rpcf_add_switch
func x1151(c conn, p Dict) Dict {
	c.Call(0x1151)
	putS(c, p, "name")
	putS(c, p, "location")
	putI(c, p, "type")
	putI(c, p, "ports_count")
	putS(c, p, "remote_id")
	putA(c, p, "address")
	putS(c, p, "login")
	putS(c, p, "password")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
