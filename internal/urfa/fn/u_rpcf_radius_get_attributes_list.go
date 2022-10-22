package fn

// rpcf_radius_get_attributes_list
func x1065(c conn, _ Dict) Dict {
	c.Call(0x1065)

	return Dict{
		"objects": getMapOf(c, func() Dict {
			obj := Dict{
				"object_id":   c.GetI(),
				"object_type": c.GetI(),
				"attrs": getMapOf(c, func() Dict {
					attr := Dict{
						"vendor_type":    c.GetI(),
						"attr_code":      c.GetI(),
						"attr_data_type": c.GetI(),
					}
					switch attr["attr_data_type"] {
					case 1:
						attr["attr_data"] = c.GetI()
					case 2:
						attr["attr_data"] = c.GetS()
					case 3:
						attr["attr_data"] = c.GetA()
					default:
						attr["attr_data"] = c.GetS()
					}
					return attr
				}),
			}
			return obj
		}),
	}
}
