package fn

// rpcf_set_iptv_custom_data
func x4518(c conn, p Dict) Dict {
	c.Call(0x4518)
	putI(c, p, "slink_id")
	putS(c, p, "data")
	c.Send()

	return Dict{
		"status": c.GetI(),
	}
}
