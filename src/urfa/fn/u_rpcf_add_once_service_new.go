package fn

// rpcf_add_once_service_new
func x2116(c conn, p Dict) Dict {
	c.Call(0x2116)
	putI(c, p, "parent_id")
	putI(c, p, "tariff_id")
	putI(c, p, "service_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putD(c, p, "cost")
	putI(c, p, "drop_from_group")
	c.Send()

	return nil
}
