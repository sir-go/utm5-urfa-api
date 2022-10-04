package fn

// rpcf_get_switch_ports_usage
func x1158(c conn, p Dict) Dict {
	c.Call(0x1158)
	putI(c, p, "switch_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1158" name="rpcf_get_switch_ports_usage">
        <input>
            <integer name="switch_id" />
        </input>
        <output>
            <integer name="ports_size" />
            <for count="ports_size" from="0" name="i">
                <integer array_index="i" name="port_id" />
                <integer array_index="i" name="users_size" />
                <for count="users_size" from="0" name="ii">
                    <integer array_index="i,ii" name="user_id" />
                    <string array_index="i,ii" name="login" />
                </for>
            </for>
        </output>
    </function>


*/
