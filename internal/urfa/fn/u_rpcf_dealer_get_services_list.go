package fn

// rpcf_dealer_get_services_list
func x70000021(c conn, p Dict) Dict {
	c.Call(0x70000021)
	putI(c, p, "service_type", -1)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000021" name="rpcf_dealer_get_services_list">
      <input>
        <integer default="-1" name="service_type" />
      </input>
      <output>
        <integer name="services_count" />
        <for count="services_count" from="0" name="i">
          <integer array_index="i" name="service_id_array" />

            <set dst="service_id_array_tmp" src="service_id_array" src_index="i" />

      <if condition="ne" value="-1" variable="service_id_array_tmp">
             <string array_index="i" name="service_name_array" />
             <integer array_index="i" name="service_type_array" />
             <string array_index="i" name="service_comment_array" />

             <if condition="ne" value="0" variable="service_type">
                 <integer name="service_status" />

                <set dst="service_status_array" dst_index="i" src="service_status" />

                <if condition="eq" value="2" variable="service_status">
                  <string array_index="i" name="tariff_name_array" />
                </if>
                <if condition="ne" value="2" variable="service_status">
                    <set dst="tariff_name_array" dst_index="i" value="" />
                </if>

             </if>

             <if condition="eq" value="0" variable="service_type">
                <set dst="service_status_array" dst_index="i" value="0" />
                <set dst="tariff_name_array" dst_index="i" value="" />
             </if>

          </if>

	  <if condition="eq" value="-1" variable="service_id_array_tmp">
             <set dst="service_name_array" dst_index="i" value="" />
             <set dst="service_type_array" dst_index="i" value="0" />
             <set dst="service_comment_array" dst_index="i" value="" />
             <set dst="service_status_array" dst_index="i" value="0" />
             <set dst="tariff_name_array" dst_index="i" value="" />
          </if>

        </for>
      </output>
    </function>


*/
