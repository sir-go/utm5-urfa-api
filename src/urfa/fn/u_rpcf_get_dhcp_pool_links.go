package fn

// rpcf_get_dhcp_pool_links
func x1160(c conn, p Dict) Dict {
	c.Call(0x1160)
	putI(c, p, "switch_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1160" name="rpcf_get_dhcp_pool_links">
        <input>
            <integer name="switch_id" />
        </input>
        <output>
            <integer name="dhcp_pools_number" />
            <for count="dhcp_pools_number" from="0" name="i">
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
                <for count="ranges_size" from="0" name="ii">
                    <ip_address array_index="ii" name="range_first_addr" />
                    <ip_address array_index="ii" name="range_last_addr" />
                </for>
            </for>
        </output>
    </function>


*/
