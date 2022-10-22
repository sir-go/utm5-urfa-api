package fn

// rpcf_dealer_get_telephony_service
func x70000028(c conn, p Dict) Dict {
	c.Call(0x70000028)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000028" name="rpcf_dealer_get_telephony_service">
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
	<double name="cost" />
	<integer name="discount_method" />
	<integer name="start_date" />
	<integer name="expire_date" />
	<integer name="radius_sessions_limit" />
	<integer name="count" />
	<for count="count" from="0" name="i">
	  <integer array_index="i" name="directions" />
	  <integer name="borders_count" />
	  <set dst="borders_count_array" dst_index="i" src="borders_count" />
	  <for count="borders_count" from="0" name="j">
	    <long array_index="i,j" name="tarif_quantity" />
	    <double array_index="i,j" name="cost" />
	  </for>
        <integer name="timerange_count" />
      <set dst="timerange_count_array" dst_index="i" src="timerange_count" />
	  <for count="timerange_count" from="0" name="j">
	    <integer array_index="i,j" name="timerange_id" />
	    <double array_index="i,j" name="cost" />
	  </for>
	</for>
	<long name="free_time" />
	<long name="first_interval" />
	<long name="first_interval_around" />
	<long name="incremental_interval" />
	<long name="unit_size" />
	<double name="fixed_call_cost" />
	<integer name="tariff_id" />
	<integer name="parent_id" />
      </output>
    </function>


*/
