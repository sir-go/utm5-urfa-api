package fn

// rpcf_add_periodic_service_batch
func x2120(c conn, p Dict) Dict {
	c.Call(0x2120)
	putI(c, p, "fictive")
	putI(c, p, "tariff_id")
	putI(c, p, "service_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putD(c, p, "cost")
	putI(c, p, "parameter")
	putI(c, p, "discount_method_t")
	putI(c, p, "start_date")
	putI(c, p, "expire_date")
	c.Send()

	return Dict{
		"service_id": c.GetI(),
	}
}
