package fn

// rpcf_get_iptraffic_service_ex
func x2234(c conn, p Dict) Dict {
	c.Call(0x2234)
	putI(c, p, "sid")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2234" name="rpcf_get_iptraffic_service_ex">
     <input>
       <integer name="sid" />
     </input>
     <output>
       <string name="service_name" />
       <string name="comment" />
       <integer name="link_by_default" />
       <integer name="is_dynamic" />
       <integer name="contract_type" />
       <double name="cost" />
       <integer name="deprecated" />
       <integer name="discount_method" />
       <integer name="sessions_limit" />
       <integer name="null_service_prepaid" />
       <integer name="borders_count" />
       <for count="borders_count" from="0" name="i">
         <integer array_index="i" name="tclassBorder" />
         <if condition="ne" value="-1" variable="tclassBorder">
           <long name="borders_size" />
           <set dst="borders_size_array" dst_index="i" src="borders_size" />
           <for count="borders_size" from="0" name="j">
             <long array_index="i,j" name="border_id" />
             <double array_index="i,j" name="border_cost" />
           </for>
         </if>
       </for>
       <integer name="prepaid_count" />
       <for count="prepaid_count" from="0" name="i">
         <integer array_index="i" name="tclassPrepaid" />
         <if condition="ne" value="-1" variable="tclassPrepaid">
           <long array_index="i" name="prepaid_amount" />
           <long array_index="i" name="prepaid_max" />
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
