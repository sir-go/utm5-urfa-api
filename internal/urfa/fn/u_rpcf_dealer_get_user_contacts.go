package fn

// rpcf_dealer_get_user_contacts
func x70000009(c conn, p Dict) Dict {
	c.Call(0x70000009)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000009" name="rpcf_dealer_get_user_contacts">
      <input>
        <integer name="user_id" />
      </input>
      <output>
          <integer name="size" />
	  <if condition="eq" value="0" variable="size">
            <error code="10" comment="user has no contacts or you dont have enough privileges" />
          </if>
	  <for count="size" from="0" name="i">
	     <integer array_index="i" name="id" />
	     <string array_index="i" name="person" />
	     <string array_index="i" name="descr" />
	     <string array_index="i" name="contact" />
	     <string array_index="i" name="email" />
	     <integer array_index="i" name="email_notify" />
	     <string array_index="i" name="short_name" />
	     <string array_index="i" name="birthday" />
	     <integer array_index="i" name="id_exec_man" />
	  </for>
      </output>
    </function>


*/
