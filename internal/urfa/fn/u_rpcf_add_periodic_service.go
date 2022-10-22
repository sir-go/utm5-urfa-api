package fn

// rpcf_add_periodic_service
func x2106(c conn, p Dict) Dict {
	c.Call(0x2106)
	putI(c, p, "fictive")
	putI(c, p, "tariff_id")
	putI(c, p, "service_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putD(c, p, "cost")
	putI(c, p, "param1")
	putI(c, p, "discount_method")
	putI(c, p, "start_date")
	putI(c, p, "expire_date")
	c.Send()

	return nil
}
