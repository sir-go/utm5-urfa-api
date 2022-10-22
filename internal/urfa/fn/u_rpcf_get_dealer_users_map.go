package fn

// rpcf_get_dealer_users_map
func x13027(c conn, p Dict) Dict {
	c.Call(0x13027)
	putI(c, p, "from")
	putI(c, p, "to")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Returns list of users with dealer link" id="0x13027" name="rpcf_get_dealer_users_map">
      <input>
        <integer name="from" />
        <integer name="to" />
      </input>
      <output>
        <integer name="users_count" />
        <for count="users_count" from="0" name="i">
         <integer array_index="i" name="user_id" />
         <string array_index="i" name="login" />
         <string array_index="i" name="full_name" />
         <integer array_index="i" name="basic_account" />
         <integer array_index="i" name="dealer_id" />
        </for>
      </output>
    </function>




*/
