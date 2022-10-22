package fn

// rpcf_get_iptv_custom_data
func x4517(c conn, p Dict) Dict {
	c.Call(0x4517)
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"data": c.GetS(),
	}
}
