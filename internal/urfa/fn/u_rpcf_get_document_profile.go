package fn

// rpcf_get_document_profile
func x460d(c conn, p Dict) Dict {
	c.Call(0x460d)
	putI(c, p, "profile_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x460d" name="rpcf_get_document_profile">
        <input>
            <integer name="profile_id" />
        </input>
        <output>
            <integer name="result" />
            <string name="name" />
            <integer name="created" />
            <integer name="modified" />
            <integer name="count" />
            <for count="count" from="0" name="i">
                <integer array_index="i" name="doc_type_array" />
                <integer array_index="i" name="template_id_array" />
            </for>
        </output>
    </function>


*/
