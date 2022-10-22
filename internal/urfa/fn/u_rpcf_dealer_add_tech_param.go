package fn

// rpcf_dealer_add_tech_param
func x70000073(c conn, p Dict) Dict {
	c.Call(0x70000073)
	putI(c, p, "type_id")
	putI(c, p, "slink_id")
	putS(c, p, "param")
	putI(c, p, "reg_date")
	putS(c, p, "password")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
