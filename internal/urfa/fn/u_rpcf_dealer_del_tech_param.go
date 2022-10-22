package fn

// rpcf_dealer_del_tech_param
func x70000076(c conn, p Dict) Dict {
	c.Call(0x70000076)
	putI(c, p, "tp_id")
	putI(c, p, "slink_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
