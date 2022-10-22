package fn

// rpcf_dealer_get_dialup_service
func x70000045(c conn, p Dict) Dict {
	c.Call(0x70000045)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000045" name="rpcf_dealer_get_dialup_service">
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
	 <string name="pool_name" />
	 <integer name="max_timeout" />
	 <integer name="null_service_prepaid" />
	 <integer name="radius_sessions_limit" />
	 <string name="login_prefix" />
	 <integer name="cost_size" />
         <for count="cost_size" from="0" name="i">
	   <string array_index="i" name="tr_name" />
	   <double array_index="i" name="param" />
	   <integer array_index="i" name="id" />
         </for>

	 <integer name="is_parent_id" />
	 <integer name="tariff_id" />
	 <integer name="parent_id" />
      </output>
    </function>


*/
