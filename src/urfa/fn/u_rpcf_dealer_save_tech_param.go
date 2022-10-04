package fn

// rpcf_dealer_save_tech_param
func x70000074(c conn, p Dict) Dict {
	c.Call(0x70000074)
	putI(c, p, "type_id")
	putI(c, p, "slink_id")
	putI(c, p, "id")
	putS(c, p, "param")
	putI(c, p, "reg_date")
	putS(c, p, "password")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
