package fn

// rpcf_get_all_services_for_user
func x2700(c conn, p Dict) Dict {
	c.Call(0x2700)
	putI(c, p, "account_id")
	c.Send()

	return Dict{
		"slinks": getMapOf(c, func() Dict {
			res := Dict{"service_id": c.GetI()}
			if res["service_id"] != -1 {
				res["service_type"] = c.GetI()
				res["service_name"] = c.GetS()
				res["tariff_name"] = c.GetS()
				res["service_cost"] = c.GetD()
				res["slink_id"] = c.GetI()
				res["discount_period_id"] = c.GetI()
			} else {
				res["service_type"] = -1
				res["service_name"] = ""
				res["tariff_name"] = ""
				res["service_cost"] = -1
				res["slink_id"] = -1
				res["discount_period_id"] = -1
			}
			return res
		}),
	}
}

/* xml:
<function id="0x2700" name="rpcf_get_all_services_for_user">
    <input>
      <integer name="account_id" />
    </input>
    <output>
      <integer name="slink_id_count" />
      <for count="slink_id_count" from="0" name="i">
        <integer name="service_id" />
        <if condition="ne" value="-1" variable="service_id">
          <set dst="service_id_array" dst_index="i" src="service_id" />
          <integer array_index="i" name="service_type_array" />
          <string array_index="i" name="service_name_array" />
          <string array_index="i" name="tariff_name_array" />
          <double array_index="i" name="service_cost_array" />
          <integer array_index="i" name="slink_id_array" />
          <integer array_index="i" name="discount_period_id_array" />
        </if>
        <if condition="eq" value="-1" variable="service_id">
          <set dst="service_id_array" dst_index="i" value="-1" />
          <set dst="service_type_array" dst_index="i" value="-1" />
          <set dst="service_name_array" dst_index="i" value="" />
          <set dst="tariff_name_array" dst_index="i" value="" />
          <set dst="service_cost_array" dst_index="i" value="-1" />
          <set dst="slink_id_array" dst_index="i" value="-1" />
          <set dst="discount_period_id_array" dst_index="i" value="-1" />
        </if>
      </for>
    </output>
  </function>


*/
