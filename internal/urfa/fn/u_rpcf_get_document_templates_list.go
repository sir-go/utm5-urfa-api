package fn

// rpcf_get_document_templates_list
func x4606(c conn, _ Dict) Dict {
	c.Call(0x4606)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4606" name="rpcf_get_document_templates_list">
        <input />
        <output>
            <integer name="cnt" />
            <for count="cnt" from="0" name="i">
                <integer array_index="i" name="template_id_array" />
                <integer array_index="i" name="doc_type_array" />
                <integer array_index="i" name="created_array" />
                <integer array_index="i" name="modified_array" />
                <string array_index="i" name="name_array" />
                <string array_index="i" name="path_array" />
            </for>
        </output>
    </function>


*/
