package fn

// rpcf_dealer_add_user_contact
func x70000042(c conn, p Dict) Dict {
	c.Call(0x70000042)
	putI(c, p, "user_id")
	putS(c, p, "person")
	putS(c, p, "descr")
	putS(c, p, "contact")
	putS(c, p, "email")
	putI(c, p, "email_notify", 1)
	putS(c, p, "short_name")
	putS(c, p, "birthday")
	putI(c, p, "id_exec_man", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000042" name="rpcf_dealer_add_user_contact">
      <input>
             <integer name="user_id" />
	     <string name="person" />
	     <string name="descr" />
	     <string name="contact" />
	     <string name="email" />
	     <integer default="1" name="email_notify" />
	     <string name="short_name" />
	     <string name="birthday" />
	     <integer default="0" name="id_exec_man" />
      </input>
      <output>
        <integer name="result" />
        <if condition="eq" value="0" variable="result">
          <error code="10" comment="user not found or you dont have enough privileges" />
        </if>
      </output>
    </function>


*/
