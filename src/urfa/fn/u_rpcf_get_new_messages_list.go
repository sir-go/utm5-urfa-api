package fn

// rpcf_get_new_messages_list
func x500f(c conn, _ Dict) Dict {
	c.Call(0x500f)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x500f" name="rpcf_get_new_messages_list">
    <input />
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
