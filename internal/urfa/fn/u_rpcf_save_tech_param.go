package fn

// rpcf_save_tech_param
func x9005(c conn, p Dict) Dict {
	c.Call(0x9005)
	putI(c, p, "type_id")
	putI(c, p, "slink_id")
	putI(c, p, "id")
	putS(c, p, "param")
	putI(c, p, "reg_date")
	putS(c, p, "passwd")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
