package fn

// rpcf_get_services_list
func x2101(c conn, p Dict) Dict {
	c.Call(0x2101)
	putI(c, p, "which_service", -1)
	c.Send()

	return Dict{
		"services": getMapOf(c, func() Dict {
			item := Dict{
				"id":      c.GetI(),
				"name":    c.GetS(),
				"type":    c.GetI(),
				"comment": c.GetS(),
			}

			if c.GetI() == 2 {
				item["tariff_name"] = c.GetS()
			} else {
				item["tariff_name"] = ""
			}

			return item
		}),
	}
}

/* xml:
<function id="0x2101" name="rpcf_get_services_list">
    <input>
      <integer default="-1" name="which_service" />
    </input>
    <output>
      <integer name="services_count" />
      <for count="services_count" from="0" name="i">
        <integer array_index="i" name="service_id_array" />
        <string array_index="i" name="service_name_array" />
        <integer array_index="i" name="service_type_array" />
        <string array_index="i" name="service_comment_array" />
        <integer array_index="i" name="service_status_array" />
        <set dst="service_status" src="service_status_array" src_index="i" />
        <if condition="eq" value="2" variable="service_status">
          <string array_index="i" name="tariff_name_array" />
        </if>
        <if condition="ne" value="2" variable="service_status">
          <set dst="tariff_name_array" dst_index="i" value="" />
        </if>
      </for>
    </output>
  </function>


*/
