package fn

// rpcf_get_ippool
func x1066(c conn, p Dict) Dict {
	c.Call(0x1066)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":      c.GetI(),
		"name":    c.GetS(),
		"address": c.GetA(),
		"mask":    c.GetI(),
	}
}
