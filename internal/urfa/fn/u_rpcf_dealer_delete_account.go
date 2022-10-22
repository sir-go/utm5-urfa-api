package fn

// rpcf_dealer_delete_account
func x70000012(c conn, p Dict) Dict {
	c.Call(0x70000012)
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000012" name="rpcf_dealer_delete_account">
      <input>
        <integer name="account_id" />
      </input>
      <output>
          <integer name="ret_code" />
          <if condition="eq" value="0" variable="ret_code">
            <error code="11" comment="account does not exist or you dont have enough privileges" />
          </if>
      </output>
    </function>


*/
