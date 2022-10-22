package fn

// rpcf_set_traffic_quota_on_link
func x5033(c conn, p Dict) Dict {
	c.Call(0x5033)
	putI(c, p, "slink_id")
	putI(c, p, "tclass_id")
	putL(c, p, "quota")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
