package fn

// rpcf_dealer_block_account
func x70000014(c conn, p Dict) Dict {
	c.Call(0x70000014)
	putI(c, p, "account_id")
	putI(c, p, "is_blocked")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000014" name="rpcf_dealer_block_account">
      <input>
	     <integer name="account_id" />
	     <integer name="is_blocked" />
      </input>
      <output>
         <integer name="aid" />
         <if condition="eq" value="0" variable="aid">
           <error code="11" comment="account not found or you dont have enough privileges" />
         </if>
      </output>
    </function>


*/
