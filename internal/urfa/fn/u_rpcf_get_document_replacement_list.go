package fn

// rpcf_get_document_replacement_list
func x460a(c conn, _ Dict) Dict {
	c.Call(0x460a)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x460a" name="rpcf_get_document_replacement_list">
        <input />
        <output>
            <integer name="rcnt" />
            <for count="rcnt" from="0" name="i">
                <string array_index="i" name="key_array" />
                <string array_index="i" name="value_array" />
            </for>
        </output>
    </function>


*/
