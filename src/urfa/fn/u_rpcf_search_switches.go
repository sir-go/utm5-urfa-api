package fn

// rpcf_search_switches
func x1156(c conn, p Dict) Dict {
	c.Call(0x1156)

	// fixme: input has a complex param
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x1156" name="rpcf_search_switches">
        <input>
            <integer name="status" />
            <integer name="params_cnt" />
            <for count="params_cnt" from="0" name="i">
                <integer array_index="i" name="field" />
                <integer array_index="i" name="criteria" />
                <string array_index="i" name="value" />
            </for>
        </input>
        <output>
            <integer name="size" />
            <if condition="ne" value="-1" variable="size">
              <if condition="ne" value="-2" variable="size">
                <if condition="ne" value="-3" variable="size">
                  <if condition="ne" value="-4" variable="size">
                    <for count="size" from="0" name="i">
                        <integer array_index="i" name="id" />
                        <string array_index="i" name="name" />
                        <string array_index="i" name="location" />
                        <integer array_index="i" name="type" />
                        <integer array_index="i" name="ports_count" />
                        <string array_index="i" name="remote_id" />
                        <ip_address array_index="i" name="address" />
                        <string array_index="i" name="login" />
                        <string array_index="i" name="password" />
                    </for>
                  </if>
                </if>
              </if>
            </if>
        </output>
    </function>


*/
