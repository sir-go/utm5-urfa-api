package fn

// rpcf_dealer_get_all_services_for_user
func x70000029(c conn, p Dict) Dict {
	c.Call(0x70000029)
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000029" name="rpcf_dealer_get_all_services_for_user">
      <input>
        <integer name="account_id" />
      </input>
      <output>
        <integer name="slink_id_count" />
        <if condition="eq" value="0" variable="slink_id_count">
             <error code="11" comment="services not found or you dont have enough privileges" />
        </if>
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
