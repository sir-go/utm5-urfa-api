package fn

// rpcf_dealer_delete_user_contact
func x70000044(c conn, p Dict) Dict {
	c.Call(0x70000044)
	putI(c, p, "user_id")
	putI(c, p, "id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000044" name="rpcf_dealer_delete_user_contact">
      <input>
             <integer name="user_id" />
	     <integer name="id" />
      </input>
      <output>
        <integer name="result" />
        <if condition="eq" value="0" variable="result">
          <error code="10" comment="contact not found or you dont have enough privileges" />
        </if>
      </output>
    </function>


*/
