package fn

// rpcf_add_contract_type
func x4802(c conn, p Dict) Dict {
	c.Call(0x4802)
	putI(c, p, "id")
	putS(c, p, "name")
	putI(c, p, "create_date")
	putI(c, p, "change_date")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
