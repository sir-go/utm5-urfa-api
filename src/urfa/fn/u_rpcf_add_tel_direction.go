package fn

// rpcf_add_tel_direction
func x10300(c conn, p Dict) Dict {
	c.Call(0x10300)
	putI(c, p, "zone_id", 0)
	putS(c, p, "name")
	putS(c, p, "calling_prefix")
	putS(c, p, "called_prefix")
	putS(c, p, "incoming_trunk")
	putS(c, p, "outgoing_trunk")
	putS(c, p, "pbx_id")
	putI(c, p, "calling_prefix_regexp", 1)
	putI(c, p, "called_prefix_regexp", 1)
	putI(c, p, "skip", 0)
	putI(c, p, "dir_type")
	c.Send()

	return Dict{
		"dir_id": c.GetI(),
	}
}
