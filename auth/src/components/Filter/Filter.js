import "./Filter.css"
import {useState} from 'react';
import {ReactComponent as Plus} from "../assets/Logos/plus.svg"
import {ReactComponent as Minus} from "../assets/Logos/minus.svg"

const Filter = (props) =>{
    const [value, setValue] = useState(<Plus className="plus"/>)
    const [open, setOpen] = useState(false)
    const [selected, setSelected] = useState(null)

    const changeButton = () =>{
        setOpen((prevOpen) => !prevOpen)
        if (open){
            setValue(<Plus className="plus"/>)
        }else{
            setValue(<Minus className="minus"/>)
        }
    }
    
    const changeAvail = (index) => {
        setSelected((prevSelected) => (prevSelected === index ? null : index));

       
        console.log(props.values[index])
    };
    
    return(
        <div className="filter-wrapper">
            <div className="filter-button">
                <h2 className="filter-name">{props.nameFilter}</h2>
                <div className="filter-btn" onClick = {() =>{
                    changeButton()
                }}>{value}
                </div>
            </div>
            {open &&<div className="filter-inside">
                {
                props.values.map((el, index) =>{
                    return <div className="filter-elements" key = {index}>
                        <div className="filter-element">{el}</div>
                        <div className="select-box" onClick = {() => {
                            changeAvail(index)
                        }}>
                            {selected === index  ? "X" : ""}
                        </div>
                    </div>
                })
            }
            </div>}
        </div>
    );
}

export default Filter