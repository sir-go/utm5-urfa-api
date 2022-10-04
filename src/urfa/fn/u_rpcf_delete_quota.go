package fn

// rpcf_delete_quota
func x20ff(c conn, p Dict) Dict {
	c.Call(0x20ff)
	putI(c, p, "slink_id")
	putI(c, p, "tc_id")
	c.Send()

	return Dict{
		"res": c.GetI(),
	}
}
