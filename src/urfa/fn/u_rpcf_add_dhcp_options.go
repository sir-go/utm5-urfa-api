package fn

// rpcf_add_dhcp_options
func x1107(c conn, p Dict) Dict {
	c.Call(0x1107)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x1107" name="rpcf_add_dhcp_options">
        <input>
           <integer name="owner_id" />
           <integer name="owner_type" />
           <integer name="size" />
           <for count="size" from="0" name="i">
                <integer array_index="i" name="option_id" />
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
                    <integer name="ip_array_size" />
                    <for count="ip_array_size" from="0" name="ii">
                        <ip_address array_index="i,ii" name="attr_data_ip" />
                    </for>
                </if>
           </for>
        </input>
        <output>
            <integer name="result" />
        </output>
   </function>


*/
