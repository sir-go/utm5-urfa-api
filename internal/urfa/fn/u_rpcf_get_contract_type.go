package fn

// rpcf_get_contract_type
func x4800(c conn, p Dict) Dict {
	c.Call(0x4800)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":          c.GetI(),
		"name":        c.GetS(),
		"create_date": c.GetI(),
		"change_date": c.GetI(),
	}
}
