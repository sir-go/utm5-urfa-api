package fn

// rpcf_edit_vod_service
func x4531(c conn, p Dict) Dict {
	c.Call(0x4531)
	putI(c, p, "service_id")
	putI(c, p, "parent_id")
	putI(c, p, "tariff_id")
	putS(c, p, "service_name")
	putS(c, p, "comment")
	putI(c, p, "link_by_default")
	putI(c, p, "is_dynamic")
	putI(c, p, "contract_type")
	putD(c, p, "cost")
	putI(c, p, "duration")
	putI(c, p, "tv_type")
	putI(c, p, "media_content")
	putL(c, p, "media_group")
	putS(c, p, "custom_data")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
