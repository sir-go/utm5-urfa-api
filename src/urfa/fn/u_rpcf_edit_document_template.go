package fn

// rpcf_edit_document_template
func x4603(c conn, p Dict) Dict {
	c.Call(0x4603)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4603" name="rpcf_edit_document_template">
        <input>
            <integer name="template_id" />
            <integer name="doc_type" />
            <string name="name" />
            <binary name="odt_data" />
        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
