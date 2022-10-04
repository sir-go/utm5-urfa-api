package fn

// rpcf_get_switch_type
func x505(c conn, p Dict) Dict {
	c.Call(0x505)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":                c.GetI(),
		"name":              c.GetS(),
		"supp_volumes":      c.GetS(),
		"port_start_offset": c.GetI(),
		"tec_id_type":       c.GetI(),
		"tec_id_len":        c.GetI(),
		"tec_id_disp":       c.GetI(),
		"tec_id_offset":     c.GetI(),
		"port_id_type":      c.GetI(),
		"port_id_len":       c.GetI(),
		"port_id_disp":      c.GetI(),
		"port_id_offset":    c.GetI(),
		"vlan_id_type":      c.GetI(),
		"vlan_id_len":       c.GetI(),
		"vlan_id_disp":      c.GetI(),
		"vlan_id_offset":    c.GetI(),
	}
}
