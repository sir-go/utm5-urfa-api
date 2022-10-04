package fn

// rpcf_create_document_template
func x4602(c conn, p Dict) Dict {
	c.Call(0x4602)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4602" name="rpcf_create_document_template">
        <input>
            <integer name="doc_type" />
            <string name="name" />
            <binary name="odt_data" />
        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
