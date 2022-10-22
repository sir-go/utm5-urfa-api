package fn

// rpcf_get_messages_list_new
func x500b(c conn, p Dict) Dict {
	c.Call(0x500b)
	putL(c, p, "time_start")
	putL(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x500b" name="rpcf_get_messages_list_new">
    <input>
      <long name="time_start" />
      <long name="time_end" />
    </input>
    <output>
      <integer name="message_size" />
      <for count="message_size" from="0" name="i">
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
