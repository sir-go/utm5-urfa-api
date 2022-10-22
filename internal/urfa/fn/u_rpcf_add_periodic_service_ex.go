package fn

// rpcf_add_periodic_service_ex
func x1522(c conn, p Dict) Dict {
	c.Call(0x1522)
	putI(c, p, "parent_id")
	putI(c, p, "tariff_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putI(c, p, "contract_type")
	putD(c, p, "cost")
	putI(c, p, "discount_method")
	putI(c, p, "scheme_id")
	c.Send()

	return Dict{
		"service_id": c.GetI(),
	}
}
