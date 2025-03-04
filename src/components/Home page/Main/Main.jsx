import React from 'react'
import './Main.css'
import { CgArrowLongRight } from "react-icons/cg";

export default function Main() {
  return (
    <main>
        <div className="container-greetings">
            <div className='text-greetings'>
                <span>Стремлениек к <br/> совершенству <br/> вместе с LEXUS</span>
                <button>ЗАКАЗАТЬ ОБРАТНЫЙ ЗВОНОК <CgArrowLongRight className='CgArrowLongRight'/></button>
            </div>
        </div>
    </main>
  )
}
