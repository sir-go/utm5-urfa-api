package fn

// rpcf_edit_switch_type
func x507(c conn, p Dict) Dict {
	c.Call(0x507)
	putI(c, p, "id")
	putS(c, p, "name")
	putS(c, p, "supp_volumes")
	putI(c, p, "port_start_offset")
	putI(c, p, "tec_id_type")
	putI(c, p, "tec_id_len")
	putI(c, p, "tec_id_disp")
	putI(c, p, "tec_id_offset")
	putI(c, p, "port_id_type")
	putI(c, p, "port_id_len")
	putI(c, p, "port_id_disp")
	putI(c, p, "port_id_offset")
	putI(c, p, "vlan_id_type")
	putI(c, p, "vlan_id_len")
	putI(c, p, "vlan_id_disp")
	putI(c, p, "vlan_id_offset")
	c.Send()

	return Dict{
		"id": c.GetI(),
	}
}
