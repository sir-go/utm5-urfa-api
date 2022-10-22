package fn

// rpcf_dealer_get_hotspot_service
func x70000024(c conn, p Dict) Dict {
	c.Call(0x70000024)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000024" name="rpcf_dealer_get_hotspot_service">
      <input>
         <integer name="service_id" />
      </input>
      <output>
          <integer name="sid" />
          <if condition="eq" value="0" variable="sid">
             <error code="11" comment="service not found or you dont have enough privileges" />
          </if>
          <string name="service_name" />
          <string name="comment" />
	  <integer name="link_by_default" />
	  <integer name="is_dynamic" />
	  <double name="cost" />
	  <integer name="deprecated" />
	  <integer name="discount_method" />
	  <integer name="start_date" />
	  <integer name="expire_date" />
	  <integer name="null_service_prepaid" />
	  <integer name="radius_sessions_limit" />
	  <double name="recv_cost" />
	  <string name="rate_limit" />
	  <integer name="hsd_allowed_net_size" />
	  <for count="hsd_allowed_net_size" from="0" name="i">
	     <integer array_index="i" name="allowed_net_id" />
	     <integer array_index="i" name="allowed_net_value" />
	  </for>
	  <integer name="cost_size" />
	  <for count="cost_size" from="0" name="i">
	     <string array_index="i" name="tr_name" />
	     <double array_index="i" name="param1" />
	     <integer array_index="i" name="param2" />
	  </for>
	  <integer name="service_data_parent_id" />
	  <integer name="tariff_id" />
	  <integer name="parent_id" />
      </output>
    </function>


*/
