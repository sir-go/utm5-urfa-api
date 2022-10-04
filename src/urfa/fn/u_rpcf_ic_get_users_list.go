package fn

// rpcf_ic_get_users_list
func x14003(c conn, p Dict) Dict {
	c.Call(0x14003)
	putI(c, p, "from")
	putI(c, p, "to")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x14003" name="rpcf_ic_get_users_list">
  	<input>
		<integer name="from" />
		<integer name="to" />
	</input>
	<output>
		<integer name="users_count" />
		<for count="users_count" from="0" name="i">
			<integer array_index="i" name="customer_id" />
			<string array_index="i" name="login" />
			<string array_index="i" name="full_name" />
			<integer array_index="i" name="status" />
			<integer array_index="i" name="last_sync_date" />
			<string array_index="i" name="ic_id" />
		</for>
	</output>
  </function>


*/
