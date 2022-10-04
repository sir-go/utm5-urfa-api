package fn

// rpcf_dealer_unlink_user_tariff
func x7000002f(c conn, p Dict) Dict {
	c.Call(0x7000002f)
	putI(c, p, "user_id")
	putI(c, p, "account_id", 0)
	putI(c, p, "tariff_link_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000002f" name="rpcf_dealer_unlink_user_tariff">
     <input>
       <integer name="user_id" />
       <integer default="0" name="account_id" />
       <integer name="tariff_link_id" />
     </input>
     <output>
        <integer name="tariff_link_id" />
        <if condition="eq" value="0" variable="tariff_link_id">
          <error code="13" comment="unable to unlink user tariff may by you dont have enough privileges" />
        </if>
     </output>
    </function>


*/
