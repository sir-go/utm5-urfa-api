package fn

// rpcf_dealer_get_switches_list
func x70000070(c conn, p Dict) Dict {
	c.Call(0x70000070)
	putI(c, p, "offset")
	putI(c, p, "count")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000070" name="rpcf_dealer_get_switches_list">
        <input>
            <integer name="offset" />
            <integer name="count" />
        </input>
        <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <string array_index="i" name="name" />
                <string array_index="i" name="location" />
                <integer array_index="i" name="type" />
                <integer array_index="i" name="ports_count" />
                <string array_index="i" name="remote_id" />
                <ip_address array_index="i" name="address" />
                <string array_index="i" name="login" />
                <string array_index="i" name="password" />
            </for>
        </output>
    </function>


*/
