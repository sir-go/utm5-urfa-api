package fn

// rpcf_user5_get_user_info
func xU4052(c conn, _ Dict) Dict {
	c.Call(-0x4052)

	return Dict{
		"user_id":           c.GetI(),
		"login":             c.GetS(),
		"basic_account":     c.GetI(),
		"balance":           c.GetD(),
		"credit":            c.GetD(),
		"is_blocked":        c.GetI(),
		"create_date":       c.GetI(),
		"last_change_date":  c.GetI(),
		"who_create":        c.GetI(),
		"who_change":        c.GetI(),
		"is_juridical":      c.GetI(),
		"full_name":         c.GetS(),
		"juridical_address": c.GetS(),
		"actual_address":    c.GetS(),
		"work_telephone":    c.GetS(),
		"home_telephone":    c.GetS(),
		"mobile_telephone":  c.GetS(),
		"web_page":          c.GetS(),
		"icq_number":        c.GetS(),
		"tax_number":        c.GetS(),
		"kpp_number":        c.GetS(),
		"bank_id":           c.GetI(),
		"bank_account":      c.GetS(),
		"int_status":        c.GetI(),
		"vat_rate":          c.GetD(),
		"passport":          c.GetS(),
		"locked_in_funds":   c.GetD(),
		"email":             c.GetS(),
	}
}
