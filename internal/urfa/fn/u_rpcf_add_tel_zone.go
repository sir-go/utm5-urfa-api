package fn

// rpcf_add_tel_zone
func x10303(c conn, p Dict) Dict {
	c.Call(0x10303)
	putS(c, p, "name")
	putI(c, p, "zone_type")
	c.Send()

	return Dict{
		"zone_id": c.GetI(),
	}
}
