package fn

// rpcf_edit_contract_type
func x4803(c conn, p Dict) Dict {
	c.Call(0x4803)
	putI(c, p, "id")
	putS(c, p, "name")
	putI(c, p, "create_date")
	putI(c, p, "change_date")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
