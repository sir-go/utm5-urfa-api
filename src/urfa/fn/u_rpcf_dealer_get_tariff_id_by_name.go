package fn

// rpcf_dealer_get_tariff_id_by_name
func x70000032(c conn, p Dict) Dict {
	c.Call(0x70000032)
	putS(c, p, "name")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000032" name="rpcf_dealer_get_tariff_id_by_name">
      <input>
	     <string name="name" />
      </input>
      <output>
	     <integer name="tid" />
              <if condition="eq" value="0" variable="tid">
               <error code="13" comment="unable to get tariff may by you dont have enough privileges" />
             </if>
      </output>
    </function>


*/
