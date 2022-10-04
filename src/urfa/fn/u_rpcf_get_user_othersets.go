package fn

// rpcf_get_user_othersets
func x9021(c conn, p Dict) Dict {
	c.Call(0x9021)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9021" name="rpcf_get_user_othersets">
        <input>
          <integer name="user_id" />
        </input>
        <output>
          <integer name="count" />
            <for count="count" from="0" name="i">
                <integer name="type" />
                <if condition="eq" value="1" variable="type">
                    <integer name="switch_id" />
                    <integer name="port" />
                </if>
                <if condition="eq" value="3" variable="type">
                    <integer name="cur_id" />
                    <string name="name" />
                </if>
            </for>
        </output>
      </function>



*/
