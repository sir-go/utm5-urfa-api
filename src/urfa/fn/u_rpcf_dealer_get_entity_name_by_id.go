package fn

// rpcf_dealer_get_entity_name_by_id
func x70000050(c conn, p Dict) Dict {
	c.Call(0x70000050)
	putI(c, p, "entity_type")
	putI(c, p, "entity_id")
	c.Send()

	return Dict{
		"entity_name": c.GetS(),
	}
}
