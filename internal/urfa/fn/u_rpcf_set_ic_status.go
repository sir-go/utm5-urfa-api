package fn

// rpcf_set_ic_status
func x14006(c conn, p Dict) Dict {
	c.Call(0x14006)
	putI(c, p, "entity_type")
	putI(c, p, "entity_id")
	putI(c, p, "ic_status")
	putS(c, p, "ic_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
