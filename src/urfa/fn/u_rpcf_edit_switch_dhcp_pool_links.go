package fn

// rpcf_edit_switch_dhcp_pool_links
func x1163(c conn, p Dict) Dict {
	c.Call(0x1163)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"error_code": c.GetI(),
	}
}

/* xml:
<function id="0x1163" name="rpcf_edit_switch_dhcp_pool_links">
        <input>
            <integer name="switch_id" />
            <integer name="number" />
            <for count="number" from="0" name="i">
                <integer name="dhcp_pool_id" />
            </for>
        </input>
        <output>
            <integer name="error_code" />
        </output>
    </function>


*/
