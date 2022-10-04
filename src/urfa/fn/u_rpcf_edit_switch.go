package fn

// rpcf_edit_switch
func x1153(c conn, p Dict) Dict {
	c.Call(0x1153)
	putI(c, p, "id")
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
