package fn

// rpcf_create_document_profile
func x4607(c conn, p Dict) Dict {
	c.Call(0x4607)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4607" name="rpcf_create_document_profile">
        <input>
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
