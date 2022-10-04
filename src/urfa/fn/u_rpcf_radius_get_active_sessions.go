package fn

// rpcf_radius_get_active_sessions
func x1073(c conn, _ Dict) Dict {
	c.Call(0x1073)

	return Dict{
		"traffic_sessions": getMapOf(c, func() Dict {
			return Dict{
				"traf_id":                    c.GetI(),
				"traf_acct_session_id":       c.GetS(),
				"traf_user_name":             c.GetS(),
				"traf_nas_ip":                c.GetA(),
				"traf_recv_date":             c.GetI(),
				"traf_last_update_date":      c.GetI(),
				"framed_ip4":                 c.GetA(),
				"framed_ip6":                 c.GetA(),
				"traffic_called_station_id":  c.GetS(),
				"traffic_calling_station_id": c.GetS(),
			}
		}),
		"tel_sessions": getMapOf(c, func() Dict {
			return Dict{
				"tel_id":                 c.GetI(),
				"tel_acct_session_id":    c.GetS(),
				"tel_user_name":          c.GetS(),
				"tel_nas_ip":             c.GetA(),
				"tel_recv_date":          c.GetI(),
				"tel_last_update_date":   c.GetI(),
				"tel_called_station_id":  c.GetS(),
				"tel_calling_station_id": c.GetS(),
			}
		}),
	}
}
