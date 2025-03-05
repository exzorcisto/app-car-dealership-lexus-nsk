import React from 'react'
import './ButtonCustom.css'

function ButtonCustom(props) {
  return <>
      <button {...props} className={'btn-custom ' + (props.className ?? '')}>
          {props.children}
      </button>
  </>
}

export default ButtonCustom;
