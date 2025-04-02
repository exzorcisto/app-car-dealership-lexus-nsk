import React from 'react'
import { Link } from 'react-router-dom';
import './NavBar.css'

function NavBar(props) {
  return <>
      <div {...props} className='nav-bar'>
          <Link to='/'>Главная</Link> \ {props.children}
      </div>
  </>
}

export default NavBar;