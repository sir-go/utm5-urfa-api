package fn

// rpcf_add_ip_to_slink
func x1510d(c conn, p Dict) Dict {
	c.Call(0x1510d)
	putI(c, p, "slink_id")
	putA(c, p, "ip")
	putI(c, p, "mask")
	putS(c, p, "mac")
	putS(c, p, "login")
	putS(c, p, "allowed_cid")
	putS(c, p, "password")
	putS(c, p, "pool_name")
	putI(c, p, "is_skip_radius", 0)
	putI(c, p, "is_skip_rfw", 0)
	putI(c, p, "router_id", 0)
	putI(c, p, "switch_id", 0)
	putI(c, p, "port_id", 0)
	putI(c, p, "vlan_id", 0)
	putI(c, p, "pool_id", 0)
	c.Send()

	return Dict{
		"result":   c.GetI(),
		"err_desc": c.GetS(),
	}
}
