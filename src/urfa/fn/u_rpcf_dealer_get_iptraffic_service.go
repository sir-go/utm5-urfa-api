package fn

// rpcf_dealer_get_iptraffic_service
func x70000025(c conn, p Dict) Dict {
	c.Call(0x70000025)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000025" name="rpcf_dealer_get_iptraffic_service">
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
	  <integer name="borders_count" />
	  <for count="borders_count" from="0" name="i">
	     <integer array_index="i" name="tclass" />

          <set dst="tclass_tmp" src="tclass" src_index="i" />

         <if condition="ne" value="-1" variable="tclass_tmp">
	       <long name="borders_size" />
	       <set dst="borders_size_array" dst_index="i" src="borders_size" />
	       <for count="borders_size" from="0" name="j">
	           <long array_index="i,j" name="border_id" />
		   <double array_index="i,j" name="border_cost" />
	       </for>
	     </if>
	     <if condition="eq" value="-1" variable="tclass_tmp">
	       <set dst="borders_size_array" dst_index="borders_size" value="0" />
	     </if>
	  </for>
	  <integer name="prepaid_count" />
	  <for count="prepaid_count" from="0" name="i">
	     <integer array_index="i" name="tclass" />

          <set dst="tclass_tmp" src="tclass" src_index="i" />

	     <if condition="ne" value="-1" variable="tclass_tmp">
	       <long array_index="i" name="prepaid_amount" />
	       <long array_index="i" name="prepaid_max" />
	     </if>
	     <if condition="eq" value="-1" variable="tclass_tmp">
	       <set dst="prepaid_amount" dst_index="i" value="0" />
	       <set dst="prepaid_max" dst_index="i" value="0" />
	     </if>
	  </for>
	  <integer name="tclass_id2group_size" />
	  <for count="tclass_id2group_size" from="0" name="i">
	       <integer array_index="i" name="tclass_id" />
   	       <integer array_index="i" name="tclass_group_id" />
	  </for>
	  <integer name="service_data_parent_id" />
	  <integer name="tariff_id" />
	  <integer name="parent_id" />
      </output>
    </function>


*/
