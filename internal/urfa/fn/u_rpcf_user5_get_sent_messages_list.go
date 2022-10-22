package fn

// rpcf_user5_get_sent_messages_list
func xU4044(c conn, p Dict) Dict {
	c.Call(-0x4044)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4044" name="rpcf_user5_get_sent_messages_list">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="message_size" />
      <for count="message_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="send_date" />
        <string array_index="i" name="subject" />
      </for>
    </output>
  </function>


*/
