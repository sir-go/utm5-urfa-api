package fn

// rpcf_get_messages_list
func x5000(c conn, p Dict) Dict {
	c.Call(0x5000)
	putL(c, p, "time_start")
	putL(c, p, "time_end")
	putI(c, p, "deprecated")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5000" name="rpcf_get_messages_list">
    <input>
      <long name="time_start" />
      <long name="time_end" />
      <integer name="deprecated" />
    </input>
    <output>
      <integer name="message_size" />
      <for count="message_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="send_date" />
        <integer array_index="i" name="recv_date" />
        <integer array_index="i" name="sender_id" />
        <string array_index="i" name="subject" />
        <string array_index="i" name="message" />
        <string array_index="i" name="mime" />
        <integer array_index="i" name="state" />
      </for>
    </output>
  </function>


*/
