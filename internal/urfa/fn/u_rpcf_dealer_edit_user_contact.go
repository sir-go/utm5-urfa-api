package fn

// rpcf_dealer_edit_user_contact
func x70000043(c conn, p Dict) Dict {
	c.Call(0x70000043)
	putI(c, p, "user_id")
	putI(c, p, "id")
	putS(c, p, "person")
	putS(c, p, "descr")
	putS(c, p, "contact")
	putS(c, p, "email")
	putI(c, p, "email_notify")
	putS(c, p, "short_name")
	putS(c, p, "birthday")
	putI(c, p, "id_exec_man")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000043" name="rpcf_dealer_edit_user_contact">
      <input>
             <integer name="user_id" />
	     <integer name="id" />
	     <string name="person" />
	     <string name="descr" />
	     <string name="contact" />
	     <string name="email" />
	     <integer name="email_notify" />
	     <string name="short_name" />
	     <string name="birthday" />
	     <integer name="id_exec_man" />
      </input>
      <output>
        <integer name="result" />
        <if condition="eq" value="0" variable="result">
          <error code="10" comment="user not found or you dont have enough privileges" />
        </if>
      </output>
    </function>


*/
