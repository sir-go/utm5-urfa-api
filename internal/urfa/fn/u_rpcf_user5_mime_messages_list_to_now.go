package fn

// rpcf_user5_mime_messages_list_to_now
func xU4033(c conn, p Dict) Dict {
	c.Call(-0x4033)
	putI(c, p, "time_start")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4033" name="rpcf_user5_mime_messages_list_to_now">
    <input>
      <integer name="time_start" />
    </input>
    <output>
      <integer name="time_end" />
      <integer name="messages_size" />
      <for count="messages_size" from="0" name="i">
        <integer array_index="i" name="send_date" />
        <integer array_index="i" name="recv_date" />
        <string array_index="i" name="subject" />
        <string array_index="i" name="message" />
        <string array_index="i" name="mime" />
        <integer array_index="i" name="state" />
      </for>
    </output>
  </function>


*/
