package fn

// rpcf_get_shaped_services
func x12006(c conn, _ Dict) Dict {
	c.Call(0x12006)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x12006" name="rpcf_get_shaped_services">
        <input />
        <output>
            <integer name="svc_count" />
            <for count="svc_count" from="0" name="i">
                <integer array_index="i" name="service_id_array" />
                <string array_index="i" name="service_name_array" />
                <string array_index="i" name="service_comment_array" />
            </for>
        </output>
    </function>


*/
