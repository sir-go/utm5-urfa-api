package fn

// rpcf_dealer_get_hotspot_services_list
func x70000022(c conn, _ Dict) Dict {
	c.Call(0x70000022)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000022" name="rpcf_dealer_get_hotspot_services_list">
      <input>
      </input>
      <output>
          <integer name="services_size" />
          <for count="services_size" from="0" name="i">
	     <integer array_index="i" name="service_id" />

              <set dst="service_id_tmp" src="service_id" src_index="i" />

         <if condition="ne" value="-1" variable="service_id_tmp">
	        <string array_index="i" name="service_name" />
	        <integer array_index="i" name="service_type" />
	        <string array_index="i" name="comment" />
	     </if>
	     <if condition="eq" value="-1" variable="service_id_tmp">
	        <set dst="service_name" dst_index="i" value="" />
	        <set dst="service_type" dst_index="i" value="0" />
	        <set dst="comment" dst_index="i" value="" />
	     </if>

          </for>
      </output>
    </function>


*/
