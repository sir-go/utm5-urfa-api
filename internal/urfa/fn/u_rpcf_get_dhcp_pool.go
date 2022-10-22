package fn

// rpcf_get_dhcp_pool
func x700(c conn, p Dict) Dict {
	c.Call(0x700)
	putI(c, p, "id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x700" name="rpcf_get_dhcp_pool">
        <input>
            <integer name="id" />
        </input>
        <output>
            <integer name="id" />
            <ip_address name="gateway" />
            <ip_address name="netmask" />
            <ip_address name="dns1_server" />
            <ip_address name="dns2_server" />
            <ip_address name="ntp_server" />
            <string name="domain_name" />
            <integer name="block_action_type" />
            <integer name="lease_time" />
            <integer name="flags" />
            <integer name="ranges_size" />
            <for count="ranges_size" from="0" name="i">
               <ip_address array_index="i" name="range_first_addr" />
               <ip_address array_index="i" name="range_last_addr" />
               <integer array_index="i" name="range_flags" />
            </for>
        </output>
    </function>


*/
