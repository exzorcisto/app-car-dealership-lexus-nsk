import React from 'react'
import './ButtonCustom.css'

function ButtonCustom(props) {
  return <>
      <button {...props} className='cstmBtn'>
          {props.children}
      </button>
  </>
}

export default ButtonCustom;
