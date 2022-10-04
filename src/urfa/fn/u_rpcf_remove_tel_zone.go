package fn

// rpcf_remove_tel_zone
func x10305(c conn, p Dict) Dict {
	c.Call(0x10305)
	putI(c, p, "zone_id")
	c.Send()

	return Dict{
		"zone_id": c.GetI(),
	}
}
