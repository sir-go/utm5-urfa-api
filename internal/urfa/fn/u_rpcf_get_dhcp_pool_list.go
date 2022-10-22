package fn

// rpcf_get_dhcp_pool_list
func x704(c conn, _ Dict) Dict {
	c.Call(0x704)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x704" name="rpcf_get_dhcp_pool_list">
        <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <ip_address array_index="i" name="gateway" />
                <ip_address array_index="i" name="netmask" />
                <ip_address array_index="i" name="dns1_server" />
                <ip_address array_index="i" name="dns2_server" />
                <ip_address array_index="i" name="ntp_server" />
                <string array_index="i" name="domain_name" />
                <integer array_index="i" name="block_action_type" />
                <integer array_index="i" name="lease_time" />
                <integer array_index="i" name="flags" />
                <integer array_index="i" name="ranges_size" />
                <for count="ranges_size" from="0" name="ii">
                    <ip_address array_index="i,ii" name="range_last_addr" />
                    <ip_address array_index="i,ii" name="range_last_addr" />
                </for>
            </for>
        </output>
    </function>


*/
