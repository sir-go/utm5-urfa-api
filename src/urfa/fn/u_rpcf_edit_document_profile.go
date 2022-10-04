package fn

// rpcf_edit_document_profile
func x4608(c conn, p Dict) Dict {
	c.Call(0x4608)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4608" name="rpcf_edit_document_profile">
        <input>
            <integer name="profile_id" />
            <string name="name" />
            <integer name="count" />
            <for count="count" from="0" name="i">
                <integer array_index="i" name="doc_type_array" />
                <integer array_index="i" name="template_id_array" />
            </for>
        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
