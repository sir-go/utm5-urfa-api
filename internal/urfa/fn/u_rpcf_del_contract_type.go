package fn

// rpcf_del_contract_type
func x4804(c conn, p Dict) Dict {
	c.Call(0x4804)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
