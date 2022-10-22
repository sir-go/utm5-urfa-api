package fn

// rpcf_get_ippools_list
func x1067(c conn, _ Dict) Dict {
	c.Call(0x1067)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1067" name="rpcf_get_ippools_list">
        <input>
        </input>
        <output>
            <integer name="size" />
            <for count="size" from="0" name="i">
              <integer array_index="i" name="id" />
              <string array_index="i" name="name" />
              <ip_address array_index="i" name="address" />
              <integer array_index="i" name="mask" />
            </for>
        </output>
    </function>


*/
