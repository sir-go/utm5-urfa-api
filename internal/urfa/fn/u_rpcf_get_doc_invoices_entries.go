package fn

// rpcf_get_doc_invoices_entries
func x7040(c conn, p Dict) Dict {
	c.Call(0x7040)
	putI(c, p, "invoice_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7040" name="rpcf_get_doc_invoices_entries">
      <input>
          <integer name="invoice_id" />
      </input>
      <output>
          <integer name="size" />
          <for count="size" from="0" name="i">
              <integer array_index="i" name="id" />
              <integer array_index="i" name="invoice_id" />
              <integer array_index="i" name="accounting_period_id" />
              <integer array_index="i" name="slink_id" />
              <integer array_index="i" name="version" />
              <integer array_index="i" name="service_type" />
              <double array_index="i" name="qnt" />
              <double array_index="i" name="base_cost" />
              <double array_index="i" name="sum_cost" />
              <double array_index="i" name="tax_amount" />
              <string array_index="i" name="name" />
              <integer array_index="i" name="det_size" />
              <for count="det_size" from="0" name="j">
                   <integer array_index="i,j" name="type" />
                   <integer array_index="i,j" name="value" />
              </for>
          </for>
      </output>
  </function>


*/
