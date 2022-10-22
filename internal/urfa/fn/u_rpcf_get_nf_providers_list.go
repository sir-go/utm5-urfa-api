package fn

// rpcf_get_nf_providers_list
func x5061(c conn, _ Dict) Dict {
	c.Call(0x5061)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5061" name="rpcf_get_nf_providers_list">
    <input />
    <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
              <integer array_index="i" name="id" />
              <integer array_index="i" name="collector_id" />
              <string array_index="i" name="name" />
              <string array_index="i" name="comments" />
              <ip_address array_index="i" name="address" />
            </for>
    </output>
  </function>


*/
