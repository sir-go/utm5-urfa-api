package fn

// rpcf_edit_dhcp_pool
func x703(c conn, p Dict) Dict {
	c.Call(0x703)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x703" name="rpcf_edit_dhcp_pool">
        <input>
            <integer name="id" />
            <ip_address name="gateway" />
            <ip_address name="netmask" />
            <ip_address name="dns1_server" />
            <ip_address name="dns2_server" />
            <ip_address name="ntp_server" />
            <string name="domain_name" />
            <integer name="block_pool_id" />
            <integer name="lease_time" />
            <integer name="flags" />
            <integer name="ranges_size" />
            <for count="ranges_size" from="0" name="i">
                <ip_address array_index="i" name="range_first_addr" />
                <ip_address array_index="i" name="range_last_addr" />
                <integer array_index="i" name="range_flags" />
            </for>
        </input>
        <output>
            <integer name="id" />
        </output>
    </function>


*/
