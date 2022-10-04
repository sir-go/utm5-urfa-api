package fn

// rpcf_get_contract_list
func x4611(c conn, p Dict) Dict {
	c.Call(0x4611)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4611" name="rpcf_get_contract_list">
        <input>
            <integer name="user_id" />
        </input>
        <output>
            <integer name="cnt" />
            <for count="cnt" from="0" name="i">
                <integer array_index="i" name="contract_id_array" />
                <integer array_index="i" name="template_id_array" />
                <integer array_index="i" name="is_old_array" />
                <string array_index="i" name="name_array" />
                <string array_index="i" name="path_array" />
            </for>
        </output>
    </function>


*/
