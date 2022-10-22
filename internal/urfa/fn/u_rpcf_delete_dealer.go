package fn

// rpcf_delete_dealer
func x13003(c conn, p Dict) Dict {
	c.Call(0x13003)
	putI(c, p, "dealer_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function comment="Delete dealer stuff" id="0x13003" name="rpcf_delete_dealer">
      <input>
        <integer name="dealer_id" />
      </input>
      <output>
          <integer name="result" />
          <if condition="eq" value="0" variable="result">
            <error code="10" comment="error deleting dealer. look debug.log for details" />
        </if>
      </output>
    </function>


*/
