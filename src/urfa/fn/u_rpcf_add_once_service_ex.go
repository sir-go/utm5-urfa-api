package fn

// rpcf_add_once_service_ex
func x1500(c conn, p Dict) Dict {
	c.Call(0x1500)
	putI(c, p, "parent_id")
	putI(c, p, "tariff_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putI(c, p, "contract_type")
	putD(c, p, "cost")
	putI(c, p, "drop_from_group")
	c.Send()

	return Dict{
		"service_id": c.GetI(),
	}
}
