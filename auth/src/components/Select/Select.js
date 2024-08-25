import "./Select.css"
import {useState} from 'react'

const Select = (props) =>{
    const [value, setValue] = useState("Sort by");
    const [open, setOpen] = useState(false);


    const changeValue = (element) =>{
        setValue(element);
        setOpen(false)
    }
    
    const hideSelect = () =>{
        setOpen((prevOpen) => !prevOpen)
    }
    
    return(
        <div className="wrapper">
            <div className = "button-sort" onClick={() =>{
                hideSelect()
            }}>{value}</div>
            {open && <div className = "select-inside">
                {
                    props.values.map((el)=> {
                        return <div onClick={ () =>{
                            changeValue(el)}}>{el}</div>
                    })
                }
            </div>
}
        </div>
    )
}

export default Select