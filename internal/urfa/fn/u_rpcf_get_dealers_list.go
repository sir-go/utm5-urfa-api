package fn

// rpcf_get_dealers_list
func x13019(c conn, _ Dict) Dict {
	c.Call(0x13019)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Returns dealer users list" id="0x13019" name="rpcf_get_dealers_list">
      <input />
      <output>
        <integer name="dealers_count" />
	<for count="dealers_count" from="0" name="i">
	    <integer array_index="i" name="dealer_id_array" />
	    <string array_index="i" name="login_array" />
	    <string array_index="i" name="full_name_array" />
	    <ip_address array_index="i" name="ip4_array" />
	    <integer array_index="i" name="mask4_array" />
	    <ip_address array_index="i" name="ip6_array" />
	    <integer array_index="i" name="mask6_array" />
	</for>
      </output>
    </function>


*/
