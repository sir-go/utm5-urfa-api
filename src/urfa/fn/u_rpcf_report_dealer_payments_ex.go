package fn

// rpcf_report_dealer_payments_ex
func x13010(c conn, p Dict) Dict {
	c.Call(0x13010)
	putI(c, p, "dealer_id")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Extended payments report" id="0x13010" name="rpcf_report_dealer_payments_ex">
      <input>
        <integer name="dealer_id" />
	<integer name="time_start" />
	<integer name="time_end" />
      </input>
      <output>
	 <integer name="methods_size" />
	 <for count="methods_size" from="0" name="i">
	   <integer array_index="i" name="method_array" />
	   <integer array_index="i" name="payments_count_array" />
	   <double array_index="i" name="payments_sum_array" />
	 </for>
      </output>
    </function>


*/
