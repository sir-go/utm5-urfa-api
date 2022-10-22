package fn

// rpcf_is_tel_number_used
func x10311(c conn, p Dict) Dict {
	c.Call(0x10311)
	putI(c, p, "slink_id")
	putS(c, p, "login")
	putS(c, p, "incoming_trunk")
	putS(c, p, "outgoing_trunk")
	putS(c, p, "pbx_id")
	putS(c, p, "number")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
