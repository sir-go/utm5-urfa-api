package fn

// rpcf_dealer_get_tariffs_list
func x7000007a(c conn, _ Dict) Dict {
	c.Call(0x7000007a)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000007a" name="rpcf_dealer_get_tariffs_list">
      <input>
      </input>
      <output>
	     <integer name="tariffs_count" />
	     <for count="tariffs_count" from="0" name="i">
	          <integer array_index="i" name="id" />
	          <string array_index="i" name="name" />
		  <integer array_index="i" name="create_date" />
		  <integer array_index="i" name="who_create" />
		  <string array_index="i" name="login" />
		  <integer array_index="i" name="change_create" />
		  <integer array_index="i" name="who_change" />
		  <string array_index="i" name="login_change" />
		  <integer array_index="i" name="expire_date" />
		  <integer array_index="i" name="is_blocked" />
		  <integer array_index="i" name="balance_rollover" />
                  <string array_index="i" name="comment" />
	     </for>
      </output>
    </function>


*/
