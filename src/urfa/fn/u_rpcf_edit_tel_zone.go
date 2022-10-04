package fn

// rpcf_edit_tel_zone
func x10304(c conn, p Dict) Dict {
	c.Call(0x10304)
	putI(c, p, "zone_id")
	putS(c, p, "name")
	putI(c, p, "zone_type")
	c.Send()

	return Dict{
		"zone_id": c.GetI(),
	}
}
