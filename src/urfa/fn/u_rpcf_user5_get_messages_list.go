package fn

// rpcf_user5_get_messages_list
func xU4043(c conn, p Dict) Dict {
	c.Call(-0x4043)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4043" name="rpcf_user5_get_messages_list">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="messages_size" />
      <for count="messages_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="send_date" />
        <integer array_index="i" name="sender_id" />
        <string array_index="i" name="subject" />
        <string array_index="i" name="mime" />
        <integer array_index="i" name="flag" />
      </for>
    </output>
  </function>


*/
