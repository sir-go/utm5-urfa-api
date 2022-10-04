package fn

// rpcf_upload_document
func x4621(c conn, p Dict) Dict {
	c.Call(0x4621)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x4621" name="rpcf_upload_document">
        <input>
            <integer name="user_id" />
            <integer name="doc_type" />
            <integer name="base_id" />
            <binary name="odt_data" />
        </input>
        <output>
            <integer name="result" />
        </output>
    </function>


*/
