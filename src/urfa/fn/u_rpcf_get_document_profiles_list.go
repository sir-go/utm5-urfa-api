package fn

// rpcf_get_document_profiles_list
func x460e(c conn, _ Dict) Dict {
	c.Call(0x460e)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x460e" name="rpcf_get_document_profiles_list">
        <input />
        <output>
            <integer name="cnt" />
            <for count="cnt" from="0" name="i">
                <integer array_index="i" name="profile_id_array" />
                <string array_index="i" name="name_array" />
                <integer array_index="i" name="created_array" />
                <integer array_index="i" name="modified_array" />
            </for>
        </output>
    </function>


*/
