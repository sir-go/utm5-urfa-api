package fn

// rpcf_get_tariff_new
func x3040(c conn, p Dict) Dict {
	c.Call(0x3040)
	putI(c, p, "tariff_id")
	c.Send()

	return Dict{
		"tariff_name":             c.GetS(),
		"tariff_create_date":      c.GetI(),
		"who_create":              c.GetI(),
		"who_create_login":        c.GetS(),
		"tariff_change_date":      c.GetI(),
		"who_change":              c.GetI(),
		"who_change_login":        c.GetS(),
		"tariff_balance_rollover": c.GetI(),
		"comments":                c.GetS(),

		"services": getMapOf(c, func() Dict {
			item := Dict{
				"id":              c.GetI(),
				"type":            c.GetI(),
				"name":            c.GetS(),
				"comment":         c.GetS(),
				"link_by_default": c.GetI(),
				"is_dynamic":      c.GetI(),
			}
			return item
		}),
	}
}
