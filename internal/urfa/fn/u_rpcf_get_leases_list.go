package fn

// rpcf_get_leases_list
func x1350(c conn, _ Dict) Dict {
	c.Call(0x1350)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1350" name="rpcf_get_leases_list">
        <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <ip_address array_index="i" name="ip" />
                <string array_index="i" name="mac" />
                <string array_index="i" name="server_id" />
                <string array_index="i" name="client_id" />
                <integer array_index="i" name="expired" />
                <integer array_index="i" name="updated" />
                <integer array_index="i" name="ipgroup_item_id" />
                <integer array_index="i" name="flags" />
            </for>
        </output>
    </function>


*/
