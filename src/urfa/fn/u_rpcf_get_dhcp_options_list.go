package fn

// rpcf_get_dhcp_options_list
func x1101(c conn, p Dict) Dict {
	c.Call(0x1101)
	putI(c, p, "id")
	putI(c, p, "type")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1101" name="rpcf_get_dhcp_options_list">
       <input>
           <integer name="id" />
           <integer name="type" />
       </input>
       <output>
           <integer name="size" />
           <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <integer array_index="i" name="option_id" />
                <integer array_index="i" name="owner_type" />
                <integer array_index="i" name="owner_id" />
                <integer array_index="i" name="data_type" />
                <if condition="eq" value="1" variable="data_type">
                    <integer array_index="i" name="attr_data_int" />
                </if>
                <if condition="eq" value="2" variable="data_type">
                    <string array_index="i" name="attr_data_string" />
                </if>
                <if condition="eq" value="3" variable="data_type">
                    <ip_address array_index="i" name="attr_data_ip" />
                </if>
                <if condition="eq" value="4" variable="data_type">
                    <string array_index="i" name="attr_data_hex_bin" />
                </if>
                <if condition="eq" value="5" variable="data_type">
                    <integer array_index="i" name="ip_array_size" />
                    <for count="ip_array_size" from="0" name="ii">
                        <ip_address array_index="i,ii" name="attr_data_ip" />
                    </for>
                </if>
           </for>
       </output>
   </function>


*/
