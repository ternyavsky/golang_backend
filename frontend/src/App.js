import { useEffect, useState } from 'react';
import './index.scss';
import axios from 'axios';
import Item from './Item';

function App() {
  const [tasks, setTasks] = useState([])
  const [activeTab, setActiveTab] = useState(0)
  const [descValue,setDescValue] = useState("")
  const [titleValue,setTitleValue] = useState("")


  useEffect( () => {
    axios.get("/api/task")
    .then( (res) => {
      setTasks(res.data)
      console.log(res.data)
    })
  }, [])

  const deleteAllTask = () => {
    axios.delete("/api/task")
    .then( (res) => {
      console.log(res.data)
      axios.get("/api/task")
      .then( (res) => {
        setTasks(res.data)
      })
    })
  }

  const deleteTask = (taskId) => {
    axios.delete(`/api/task/${taskId}`)
    .then( (res) => {
      console.log(res.data)
      axios.get("/api/task")
    .then( (resp) => {
      setTasks(resp.data)
      console.log(resp.data)
    })
    
  })
  }
  const addTask = (titled,descriptioned) => {
      axios.post("/api/task",{ title:titled, description: descriptioned, status:false })
    .then( (res) => {
      console.log(res.data)
            axios.get("/api/task")
 
    .then((res) => {
           setDescValue("")
      setTitleValue("")
      setTasks(res.data)
    })
  })
  }

  const putCheckBox = (el) =>{
    console.log(!el.status)
    axios.put(`/api/task/${el.id}`, {title:el.title,
     description: el.description,
     status: !el.status})
     .then( (res) => {
      axios.get("/api/task")
      
      .then( (res) => {
        setTasks(res.data)
        console.log(tasks)
      })
     })
  } 

  return (
    <div className="App">
      <div className="top">
        <ul className="tabs">
          <li onClick={ () => setActiveTab(0)} className={activeTab === 0 ? "active" : ""}>Все</li>
          <li onClick={() => setActiveTab(false)} className={activeTab === false ? "active" : ""}>Открытые </li>
          <li onClick={() => setActiveTab(true)} className={activeTab === true ? "active" : ""}>Завершённые</li>
        </ul>
        <svg height="24" viewBox="0 0 24 24" width="24" xmlns="http://www.w3.org/2000/svg" onClick={ () => deleteAllTask()}>
          <path d="M15 2H9c-1.103 0-2 .897-2 2v2H3v2h2v12c0 1.103.897 2 2 2h10c1.103 0 2-.897 2-2V8h2V6h-4V4c0-1.103-.897-2-2-2zM9 4h6v2H9V4zm8 16H7V8h10v12z" />
        </svg>
      </div>

    
    {  
    tasks.filter( (obj) => {return  activeTab === 0 ? obj : obj.status === activeTab }).map( (obj) => (
      <Item obj={obj} changeCheckBox={putCheckBox} taskDelete={deleteTask}/>
    ))} 
    <div className="form">
        <div className="form__checkbox">
          
        </div>
        <div class="form__fields">
          <input type="text" id="title" placeholder="Название"  className="input-title" value={titleValue} onChange={ (el) => setTitleValue(el.target.value) }/>
          <input type="text" id="description" placeholder="Введите текст..." className="input-text" value={descValue} onChange={ (el) => setDescValue(el.target.value) } />
        </div>
        <svg height="32px" viewBox="0 0 512 512" width="32px" onClick={
             () => {
              let title = document.getElementById("title").value;
              let description = document.getElementById("description").value
              addTask(title, description)

            }
             }>
          <g>
            <g>
              <g>
                <path d="M256,48C141.1,48,48,141.1,48,256s93.1,208,208,208c114.9,0,208-93.1,208-208S370.9,48,256,48z M256,446.7     c-105.1,0-190.7-85.5-190.7-190.7S150.9,65.3,256,65.3S446.7,150.9,446.7,256S361.1,446.7,256,446.7z" />
              </g>
            </g>
            <g>
              <polygon points="264.1,128 247.3,128 247.3,247.9 128,247.9 128,264.7 247.3,264.7 247.3,384 264.1,384 264.1,264.7 384,264.7     384,247.9 264.1,247.9   " />
            </g>
          </g>
        </svg>
      </div>
    </div>
  );
}

export default App;
