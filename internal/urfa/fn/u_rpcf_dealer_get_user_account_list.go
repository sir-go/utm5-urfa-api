package fn

// rpcf_dealer_get_user_account_list
func x70000008(c conn, p Dict) Dict {
	c.Call(0x70000008)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000008" name="rpcf_dealer_get_user_account_list">
      <input>
        <integer name="user_id" />
      </input>
      <output>
          <integer name="accounts_count" />
	  <if condition="eq" value="0" variable="accounts_count">
            <error code="10" comment="user has no accounts or you dont have enough privileges" />
          </if>
	  <for count="accounts_count" from="0" name="i">
	     <integer array_index="i" name="account" />
	     <string array_index="i" name="account_name" />
	  </for>
      </output>
    </function>


*/
