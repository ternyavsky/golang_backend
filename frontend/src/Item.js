import React from 'react'

export default function Item({obj,changeCheckBox,taskDelete}) {
  return (
       <div className="item">
 <input class="styled-checkbox" id={`checkbox${obj.id}`} checked={obj.status} onClick={ () => {
  changeCheckBox(obj);

  }} type="checkbox" value="value1" />
 <label for={`checkbox${obj.id}`} />
 <div className="item__content">
   {obj.status === false ? <b>{obj.title}</b> : <s><b>{obj.title}</b></s>}
   <p>
   {obj.status === false ? <b>{obj.description}</b> : <s><b>{obj.description}</b></s>}
   </p>
   <div className="item__bottom">
     <ul>
       <li>
         <svg style={{ width: 24, height: 24 }} viewBox="0 0 70 70">
           <path d="M51,19h-4v-4h-4v4H27v-4h-4v4h-4c-2.209,0-4,1.791-4,4v28c0,2.209,1.791,4,4,4h32c2.209,0,4-1.791,4-4V23   C55,20.791,53.209,19,51,19z M51,51H19V31h32V51z M51,27H19v-4h32V27z" />
           <rect height="4" width="4" x="35" y="35" />
           <rect height="4" width="4" x="43" y="35" />
           <rect height="4" width="4" x="35" y="43" />
           <rect height="4" width="4" x="27" y="43" />
         </svg>
         <span>{obj.date.slice(0,10)}</span>
       </li>
       <li>
         <svg viewBox="0 0 24 24" onClick={ () => taskDelete(obj.id)} xmlns="http://www.w3.org/2000/svg">
           <path d="M15 2H9c-1.103 0-2 .897-2 2v2H3v2h2v12c0 1.103.897 2 2 2h10c1.103 0 2-.897 2-2V8h2V6h-4V4c0-1.103-.897-2-2-2zM9 4h6v2H9V4zm8 16H7V8h10v12z" />
         </svg>
       </li>
       <li>
        
       </li>
     </ul>
   </div>
   </div>
</div>

  )
}
