package fn

// rpcf_report_dealer_users
func x1300e(c conn, p Dict) Dict {
	c.Call(0x1300e)
	putI(c, p, "dealer_id")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Report on created users" id="0x1300e" name="rpcf_report_dealer_users">
      <input>
        <integer name="dealer_id" />
	<integer name="time_start" />
	<integer name="time_end" />
      </input>
      <output>
        <integer name="users_count" />
	<for count="users_count" from="0" name="i">
	  <string array_index="i" name="user_id_array" />
          <string array_index="i" name="login_array" />
	  <string array_index="i" name="full_name_array" />
	  <integer array_index="i" name="create_date_array" />
	</for>
      </output>
    </function>


*/
