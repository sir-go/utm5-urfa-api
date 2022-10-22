package fn

// rpcf_dealer_get_document
func x7000007b(c conn, p Dict) Dict {
	c.Call(0x7000007b)
	putI(c, p, "user_id")
	putI(c, p, "doc_type")
	putI(c, p, "base_id")
	putI(c, p, "file_type")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000007b" name="rpcf_dealer_get_document">
        <input>
            <integer name="user_id" />
            <integer name="doc_type" />
            <integer name="base_id" />
            <integer name="file_type" />
        </input>
        <output>
            <integer name="result" />
            <binary name="doc_data" />
        </output>
    </function>


*/
