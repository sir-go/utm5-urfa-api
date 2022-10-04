package fn

// rpcf_remove_tclass
func x2304(c conn, p Dict) Dict {
	c.Call(0x2304)
	putI(c, p, "tclass_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
