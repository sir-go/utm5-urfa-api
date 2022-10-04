package fn

// rpcf_get_dealer_contacts
func x13011(c conn, p Dict) Dict {
	c.Call(0x13011)
	putI(c, p, "dealer_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Get dealer contact info" id="0x13011" name="rpcf_get_dealer_contacts">
      <input>
         <integer name="dealer_id" />
      </input>
      <output>
         <integer name="size" />
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
