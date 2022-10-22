package fn

// rpcf_add_once_service
func x210b(c conn, p Dict) Dict {
	c.Call(0x210b)
	putI(c, p, "parent_id")
	putI(c, p, "tariff_id")
	putI(c, p, "service_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putD(c, p, "cost")
	c.Send()

	return nil
}
